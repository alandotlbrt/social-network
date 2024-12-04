package profile

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"socialNetwork/api/auth"
	"socialNetwork/pkg/db"
	config "socialNetwork/pkg/db/var"
	"strings"
)

var Password string
var Pictures string

func UpdateUserInformations(w http.ResponseWriter, r *http.Request, information config.InformationRegister) {
	database, err := db.OpenDB("database")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	var Username interface{}
	var hashedPassword string

	if information.Password != "" {
		if !auth.VerifyPassword(information.Password, Password) {
			json.NewEncoder(w).Encode(config.Response{Success: false, Id: 1, Message: "Wrong MDP"})
			return
		} else {
			eror := verifiationMpd(information.NewPassword)
			if len(eror) > 0 {
				json.NewEncoder(w).Encode(config.Response{Success: false, Id: 1, Errors: eror})
				return
			}

			hashedPassword, err = auth.HashPassword(information.NewPassword)
			if err != nil {
				fmt.Print("Password not Hash ", err)
				return
			}
			Password = hashedPassword
		}
	}

	if information.Username == "" {
		Username = nil
	} else {
		Username = information.Username
	}
	if information.Photo != "" {
		Pictures = information.Photo
	}

	updateSQL := fmt.Sprintf(`UPDATE users SET (email, first_name, last_name, pp, nickname, about_me, privacy, password) = (?, ?, ?, ?, ?, ?, ?, ?) WHERE id_user = %s`, information.Id)
	_, err = database.Exec(updateSQL, information.Email, information.Firstname, information.Lastname, Pictures, Username, information.About, information.Privacy, Password)
	if err != nil {
		if strings.Contains(err.Error(), "nickname") {
			json.NewEncoder(w).Encode(config.Response{Success: false, Id: 2, Message: " Username allready taken"})
			return
		}
		if strings.Contains(err.Error(), "email") {
			json.NewEncoder(w).Encode(config.Response{Success: false, Id: 3, Message: " Email allready taken"})
			return
		}
	}
	json.NewEncoder(w).Encode(config.Response{Success: true})
}

func verifiationMpd(NewPass string) []string {
	var errors []string
	capital := `[A-Z]`
	check, _ := regexp.MatchString(capital, NewPass)
	if !check {
		errors = append(errors, "Password must contain at least one capital letter")
	}

	ree := `[a-z]`
	check, _ = regexp.MatchString(ree, NewPass)
	if !check {
		errors = append(errors, "Password must contain at least one lowercase letter")
	}

	ree = `\d`
	check, _ = regexp.MatchString(ree, NewPass)
	if !check {
		errors = append(errors, "Password must contain at least one digit")
	}

	if len(NewPass) < 8 {
		errors = append(errors, "Password must be at least 8 characters long")
	}

	ree = `[@$!%*?&._-]`
	check, _ = regexp.MatchString(ree, NewPass)
	if !check {
		errors = append(errors, "Password must contain at least one special character (@, $, !, %, *, ?, &, .,-,_ )")
	}

	return errors
}
