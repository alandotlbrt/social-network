package post

import (
	"database/sql"
	"fmt"
	"socialNetwork/pkg/db"
	config "socialNetwork/pkg/db/var"
)

func LoadPostsModel() ([]config.PostStruct, bool) {
	var rows *sql.Rows
	db, err := db.OpenDB("database")
	if err != nil {
		fmt.Println(err)
		return []config.PostStruct{}, false
	}
	rows, err = db.Query("SELECT * FROM 'post'")
	if err != nil {
		fmt.Println(err)
		return []config.PostStruct{}, false
	}
	var everyPostSelected []config.PostStruct
	for rows.Next() {
		var postSelected config.PostStruct
		rows.Scan(&postSelected.Id_post, &postSelected.Id_user, &postSelected.Content, &postSelected.Image, &postSelected.Privacy, &postSelected.Like)

		everyPostSelected = append(everyPostSelected, postSelected)
	}

	return everyPostSelected, true
}
