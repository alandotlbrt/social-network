package websocket

import (
	"database/sql"
	"fmt"
	"log"
	"socialNetwork/pkg/db"
)

var Send struct {
	Socket   interface{}
	Pictures string
}

func mapToArray(dataSet map[string]bool) []string {
	dataArray := make([]string, 0, len(dataSet))
	for key := range dataSet {
		dataArray = append(dataArray, key)
	}
	return dataArray
}
func registerEventHandler(eventType string, handler EventHandler) {
	eventHandlers[eventType] = handler
}

func init() {
	registerEventHandler("Message", Message)
	registerEventHandler("add", handleAddEvent)
	registerEventHandler("delete", handleDeleteEvent)
}

func Message(client *Client, message map[string]interface{}) {
	saveInDatabase(message)

	profil := takeUserInfoInDatabase(message["from"])
	Send.Pictures = profil
	Send.Socket = message
	for _, v := range clients {
		v.conn.WriteJSON(Send)
	}
}

func saveInDatabase(message map[string]interface{}) {
	db, err := db.OpenDB("database")
	if err != nil {
		fmt.Println("database error ")
	}

	query := `INSERT INTO chat (from_id, to_id, content, type, pictures) VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(message["from"], 0, message["value"], message["style"], message["image"])
	if err != nil {
		fmt.Println(err)
	}
}

func takeUserInfoInDatabase(id interface{}) string {

	var profil_Pictures string
	db, err := db.OpenDB("database")
	if err != nil {

	}

	err = db.QueryRow("SELECT pp FROM users WHERE id_user = ?", id).Scan(&profil_Pictures)
	if err != nil {
		if err == sql.ErrNoRows {
		}

	}
	return profil_Pictures
}
