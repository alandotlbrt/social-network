package auth

import (
	"log"
	"net/http"
	"regexp"

	"socialNetwork/pkg/db"
	config "socialNetwork/pkg/db/var"
	"strings"
	"time"
)

func RegisterFunctions(w http.ResponseWriter, r *http.Request, userInformations config.InformationRegister) string {

	err := errorRegisterHandler(userInformations)
	if err != "" {
		return err
	}

	err = registerInsertInDb(userInformations)
	if err != "" {
		return err
	}

	return ""
}

func registerInsertInDb(userInformations config.InformationRegister) string {

	db, err := db.OpenDB("database")
	if err != nil {
		return "database error "
	}
	defer db.Close()

	hashPass, err := HashPassword(userInformations.Password)
	if err != nil {
		return "error hashed password"
	}

	var Username interface{}

	if userInformations.Username == "" {
		Username = nil
	} else {
		Username = userInformations.Username
	}

	query := `INSERT INTO Users (email, password, first_name, last_name, birthday, pp, nickname, about_me, privacy, follower, follow) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(userInformations.Email, hashPass, userInformations.Firstname, userInformations.Lastname, userInformations.Birthday, userInformations.Photo, Username, userInformations.About, userInformations.Privacy, 0, 0)
	if err != nil {
		if strings.Contains(err.Error(), "nickname") {
			return "Username allready taken"
		}
		if strings.Contains(err.Error(), "email") {
			return " Email allready taken"
		}
	}
	return ""
}

func errorRegisterHandler(userInformations config.InformationRegister) string {
	registerErrors := verifiationRegister(userInformations)
	if len(registerErrors) > 0 {
		return registerErrors[0]
	}
	return ""
}

func verifiationRegister(information config.InformationRegister) []string {
	var errors []string
	ree := `^[^\s@]+@[^\s@]+\.[^\s@]+$`
	check, _ := regexp.MatchString(ree, information.Email)
	if !check {
		errors = append(errors, "Email not valid")
	}

	if information.Birthday == "" {
		errors = append(errors, "Birthday not valid")
	} else {
		var today = time.Now()
		layout := "2006-01-02"
		birthDate, err := time.Parse(layout, information.Birthday)
		if err != nil {
			errors = append(errors, "Birthday format is invalid")
		} else {
			age := today.Year() - birthDate.Year()
			monthDiff := today.Month() - birthDate.Month()
			if monthDiff < 0 || (monthDiff == 0 && today.Day() < birthDate.Day()) {
				age--
			}
			if age <= 15 {
				errors = append(errors, "Birthday indicates too young")
			}
		}
	}

	capital := `[A-Z]`
	check, _ = regexp.MatchString(capital, information.Password)
	if !check {
		errors = append(errors, "Password must contain at least one capital letter")
	}

	ree = `[a-z]`
	check, _ = regexp.MatchString(ree, information.Password)
	if !check {
		errors = append(errors, "Password must contain at least one lowercase letter")
	}

	ree = `\d`
	check, _ = regexp.MatchString(ree, information.Password)
	if !check {
		errors = append(errors, "Password must contain at least one digit")
	}

	if len(information.Password) < 8 {
		errors = append(errors, "Password must be at least 8 characters long")
	}

	ree = `[@$!%*?&._-]`
	check, _ = regexp.MatchString(ree, information.Password)
	if !check {
		errors = append(errors, "Password must contain at least one special character (@, $, !, %, *, ?, &, .,-,_ )")
	}
	if information.Password != information.ConfirmPassword {
		errors = append(errors, "Passwords do not match")
	}

	return errors
}
