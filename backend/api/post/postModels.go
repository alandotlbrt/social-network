package post

import (
	"fmt"
	"socialNetwork/pkg/db"
)

func CreatePost(idUser, message, pictures, privacy string) string {

	errorMessage := postIsValid(message, pictures, privacy)
	fmt.Println(errorMessage)
	if errorMessage != "" {
		return errorMessage
	}

	database, err := db.OpenDB("database")
	if err != nil {
		fmt.Println(err)
		return "an error has occured"
	}
	defer database.Close()

	query := `INSERT INTO post (id_user, content, image, privacy, like) VALUES (?, ?, ?, ?, ?)`
	stmt, err := database.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return "an error has occured"
	}
	defer stmt.Close()

	_, err = stmt.Exec(idUser, message, pictures, privacy, 0)
	if err != nil {
		fmt.Println(err)
		return "an error has occured"

	}

	return ""
}

func postIsValid(message, pictures, privacy string) string {

	if len(message) < 2 && pictures == "" {
		return "message is too short"
	}

	if privacy != "private" && privacy != "public" {
		return "you need to chose a privacy for your post"
	}

	if len(message) > 1000 {
		return "message is too long"
	}

	return ""
}
