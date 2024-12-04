package Chat

import (
	"database/sql"
	"fmt"
	"socialNetwork/pkg/db"
	config "socialNetwork/pkg/db/var"
)

func RetreiveChat() ([]config.ChatStruct, bool) {
	var rows *sql.Rows
	db, err := db.OpenDB("database")
	if err != nil {
		fmt.Println(err)
		return []config.ChatStruct{}, false
	}
	rows, err = db.Query("SELECT * FROM 'chat'")
	if err != nil {
		fmt.Println(err)
		return []config.ChatStruct{}, false
	}
	var everyChatSelected []config.ChatStruct
	for rows.Next() {
		var sousRows *sql.Rows
		var chatSelected config.ChatStruct
		rows.Scan(&chatSelected.Id_chat, &chatSelected.From_Id, &chatSelected.To_Id, &chatSelected.Content, &chatSelected.Type, &chatSelected.Image)
		sousRows, err = db.Query("SELECT pp FROM users WHERE id_user = ?", chatSelected.From_Id)
		if err != nil {
			fmt.Println(err)
			return []config.ChatStruct{}, false
		}
		for sousRows.Next() {
			sousRows.Scan(&chatSelected.Pictures)
		}
		everyChatSelected = append(everyChatSelected, chatSelected)
	}
	defer db.Close()
	return everyChatSelected, true

}

func RetreiveUser() ([]config.UserChatStruct, bool) {
	var rows *sql.Rows
	db, err := db.OpenDB("database")
	if err != nil {
		fmt.Println(err)
		return []config.UserChatStruct{}, false
	}
	rows, err = db.Query("SELECT id_user, nickname, last_name, first_name, pp FROM 'users'")
	if err != nil {
		fmt.Println(err)
		return []config.UserChatStruct{}, false
	}
	var everyUserSelected []config.UserChatStruct
	for rows.Next() {
		var userSelected config.UserChatStruct
		rows.Scan(&userSelected.Id_User, &userSelected.Username, &userSelected.Lastname, &userSelected.Firstname, &userSelected.Pictures)
		everyUserSelected = append(everyUserSelected, userSelected)
	}
	defer db.Close()
	return everyUserSelected, true

}
