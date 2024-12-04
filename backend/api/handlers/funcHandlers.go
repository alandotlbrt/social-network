package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	profile "socialNetwork/api/user"
	"socialNetwork/pkg/db"
	config "socialNetwork/pkg/db/var"
	"socialNetwork/sessions"
)

func validationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	_, valideSessions := sessions.ValidateSession(w, r)
	if !valideSessions {
		json.NewEncoder(w).Encode(map[string]bool{"valid": false})
	} else {
		json.NewEncoder(w).Encode(map[string]bool{"valid": true})
	}

}
func userDataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	idUser, err := db.GetInDatabase(nil, "users", "cookie", sessions.GetCookie(r), "id_user")
	if err != nil {
		fmt.Println(err)
		return
	}

	iduserWhoChecking, err := db.CheckInDatabaseAndGet(nil, "users", "cookie", sessions.GetCookie(r), "id_user")
	if err != nil {
		fmt.Println(err)
		return
	}

	validData, structUserData := profile.GetUserData(w, iduserWhoChecking, idUser)

	if validData {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(config.Response{Success: true, Info: structUserData, Message: "Login successful"})
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(config.Response{Success: false, Message: "An error has occurred"})
	}
}

func otherUserDataHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	var reqBody Register

	if err = json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		fmt.Println("Erreur de décodage JSON:", err)
		return
	}

	iduserWhoChecking, err := db.CheckInDatabaseAndGet(nil, "users", "cookie", sessions.GetCookie(r), "id_user")
	if err != nil {
		fmt.Println(err)
		return
	}

	validData, structUserData := profile.GetUserData(w, iduserWhoChecking, reqBody.Id)

	if validData {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(config.Response{Success: true, Info: structUserData, Message: "Login successful"})
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(config.Response{Success: false, Message: "An error has occurred"})
	}
}
