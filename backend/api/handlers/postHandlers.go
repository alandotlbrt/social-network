package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"socialNetwork/api/post"
	"socialNetwork/pkg/db"
	"socialNetwork/sessions"
)

type postStruct struct {
	Message  string `json:"message"`
	Pictures string `json:"pictures"`
	Privacy  string `json:"privacy"`
}

func createPostHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	var postDecoder postStruct

	err := json.NewDecoder(r.Body).Decode(&postDecoder)
	if err != nil {
		fmt.Println(err)
		return
	}

	idUser, err := db.GetInDatabase(nil, "users", "cookie", sessions.GetCookie(r), "id_user")
	if err != nil {
		fmt.Println(err)
		return
	}

	validationMessage := post.CreatePost(idUser, postDecoder.Message, postDecoder.Pictures, postDecoder.Privacy)
	if validationMessage == "" {
		json.NewEncoder(w).Encode(Response{Success: true, Message: "everything's good"})

	} else {

		json.NewEncoder(w).Encode(Response{Success: false, Message: validationMessage})
	}

}

func loadPost(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	everyPostSelected, err := post.LoadPostsModel()
	if err {
		json.NewEncoder(w).Encode(everyPostSelected)
	} else {
		json.NewEncoder(w).Encode("error")
	}
}
