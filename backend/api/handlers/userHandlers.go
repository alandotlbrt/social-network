package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	auth "socialNetwork/api/auth"
	profile "socialNetwork/api/user"
	"socialNetwork/pkg/db"
	config "socialNetwork/pkg/db/var"
	"socialNetwork/sessions"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("hdzad")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	var creds Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	validation, message := auth.LoginMainFunctions(w, r, creds.Email, creds.Password)
	if validation {
		if err = sessions.CreateSession(w, r, creds.Email); err != nil {
			fmt.Println(err)
			return
		}
		json.NewEncoder(w).Encode(config.Response{Success: true, Message: message})
	} else {
		json.NewEncoder(w).Encode(config.Response{Success: false, Message: message})

	}

}

func registerHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	var data Register

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	userRegisterInformations := config.InformationRegister{
		Username:        data.Username,
		Firstname:       data.Firstname,
		Lastname:        data.Lastname,
		Password:        data.Password,
		ConfirmPassword: data.ConfirmPassword,
		Birthday:        data.Birthday,
		Photo:           data.Photo,
		Email:           data.Email,
		About:           data.About,
		Privacy:         data.Privacy,
	}

	errorMessage := auth.RegisterFunctions(w, r, userRegisterInformations)
	if errorMessage != "" {
		json.NewEncoder(w).Encode(config.Response{Success: false, Message: errorMessage})
	} else {
		json.NewEncoder(w).Encode(config.Response{Success: true, Message: errorMessage})
	}

}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	var dataFetched Register
	err := json.NewDecoder(r.Body).Decode(&dataFetched)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	userRegisterInformations := config.InformationRegister{
		Id:          dataFetched.Id,
		Username:    dataFetched.Username,
		Firstname:   dataFetched.Firstname,
		Lastname:    dataFetched.Lastname,
		Photo:       dataFetched.Photo,
		Email:       dataFetched.Email,
		About:       dataFetched.About,
		Password:    dataFetched.Password,
		NewPassword: dataFetched.NewPassword,
		Privacy:     dataFetched.Privacy,
	}

	profile.UpdateUserInformations(w, r, userRegisterInformations)
}

func followHandlers(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	var followJSON followJSONStruct

	err := json.NewDecoder(r.Body).Decode(&followJSON)
	if err != nil {
		json.NewEncoder(w).Encode("error")
		return
	}

	if followJSON.IdToFollow == followJSON.IdUser {
		fmt.Println("error")
		json.NewEncoder(w).Encode("he is the same profile")
		return
	}

	if profile.UserFollowModels(followJSON.IdUser, followJSON.IdToFollow) {
		json.NewEncoder(w).Encode(true)
	} else {
		http.Error(w, "already follow", http.StatusNonAuthoritativeInfo)
	}
}

func unFollowHandlers(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	var followJSON followJSONStruct

	err := json.NewDecoder(r.Body).Decode(&followJSON)
	if err != nil {
		json.NewEncoder(w).Encode("error")
		return
	}

	if followJSON.IdToUNFollow == followJSON.IdUser {
		json.NewEncoder(w).Encode("he is the same profile")
		return
	}

	if profile.UserUnFollowModels(followJSON.IdUser, followJSON.IdToUNFollow) {
		json.NewEncoder(w).Encode(true)
	} else {
		http.Error(w, "already unfollow", http.StatusNonAuthoritativeInfo)
	}

}

func requestFollowHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	var followJSON followJSONStruct

	err := json.NewDecoder(r.Body).Decode(&followJSON)
	if err != nil {
		fmt.Println(err)
		return
	}

	if followJSON.IdToFollow == followJSON.IdUser {
		json.NewEncoder(w).Encode("he is the same profile")
		return
	}

	if profile.WaitingFollowModel(followJSON.IdUser, followJSON.IdToFollow) {
		json.NewEncoder(w).Encode(true)
	} else {
		http.Error(w, "already unfollow", http.StatusNonAuthoritativeInfo)
	}

}

func unRequestFollowHandlers(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	var followJSON followJSONStruct

	err := json.NewDecoder(r.Body).Decode(&followJSON)
	if err != nil {
		fmt.Println(err)
		return
	}

	if followJSON.IdToUNFollow == followJSON.IdUser {
		json.NewEncoder(w).Encode("he is the same profile")
		return
	}

	if profile.UnRequestFollowModels(followJSON.IdUser, followJSON.IdToUNFollow) {
		json.NewEncoder(w).Encode(true)
	} else {
		http.Error(w, "already unfollow", http.StatusNonAuthoritativeInfo)
	}
}

func waitingFollowersHandlers(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	iduserWhoAccept, err := db.CheckInDatabaseAndGet(nil, "users", "cookie", sessions.GetCookie(r), "id_user")
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode("an error has occured")
		return
	}

	waitingUser := profile.FetchWaitingFollowersModels(iduserWhoAccept)

	if waitingUser != nil {
		json.NewEncoder(w).Encode(waitingUser)
	} else {
		json.NewEncoder(w).Encode("an error has occured")
	}

}

func declineFollowerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	var followJSON followJSONStruct

	err := json.NewDecoder(r.Body).Decode(&followJSON)
	if err != nil {
		fmt.Println(err)
		return
	}

	if profile.UnRequestFollowModels(followJSON.IdToDecline, followJSON.IdUser) {
		json.NewEncoder(w).Encode(true)
	} else {
		json.NewEncoder(w).Encode(false)
	}

}

func acceptFollowerHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	var followJSON followJSONStruct

	err := json.NewDecoder(r.Body).Decode(&followJSON)
	if err != nil {
		fmt.Println(err)
		return
	}

	if profile.AcceptFollower(followJSON.IdToAccept, followJSON.IdUser) {
		json.NewEncoder(w).Encode(true)
	} else {
		json.NewEncoder(w).Encode(false)
	}

}

func fetchfollowersHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	var followJSON followJSONStruct

	err := json.NewDecoder(r.Body).Decode(&followJSON)
	if err != nil {
		fmt.Println(err)
		return
	}

	everyFollowers := profile.FetchFollowersFromProfile(followJSON.IdUser, followJSON.IdUserState, followJSON.Type)

	if everyFollowers != nil {
		json.NewEncoder(w).Encode(everyFollowers)
	} else {
		json.NewEncoder(w).Encode(false)
	}

}

func removeFollowerHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	var followJSON followJSONStruct

	err := json.NewDecoder(r.Body).Decode(&followJSON)
	if err != nil {
		fmt.Println(err)
		return
	}

	if profile.UserUnFollowModels(followJSON.IdToUNFollow, followJSON.IdUser) {
		json.NewEncoder(w).Encode(true)
	} else {
		json.NewEncoder(w).Encode(false)
	}
}
