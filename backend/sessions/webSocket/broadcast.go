package websocket

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)
func handleAddEvent(client *Client, message map[string]interface{}) {
	value := message["value"].(string)
	adds(dataSet, value)
}

func handleDeleteEvent(client *Client, message map[string]interface{}) {
	valueToDelete := message["value"].(string)
	deletes(dataSet, valueToDelete)
}
func deletes(dataSet map[string]bool, valueToDelete string) {
	lock.Lock()
	delete(dataSet, valueToDelete)
	lock.Unlock()
	broadcast()
}

func adds(dataSet map[string]bool, value string) {
	lock.Lock()
	dataSet[value] = true
	lock.Unlock()
	broadcast()
}
func broadcast() {
	lock.Lock()
	defer lock.Unlock()
	dataArray := mapToArray(dataSet)
	msg, err := json.Marshal(map[string]interface{}{
		"type": "update",
		"data": dataArray,
	})
	if err != nil {
		fmt.Println("Error marshalling broadcast message:", err)
		return
	}
	for conn := range clients {
		err := conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			fmt.Println("Error sending broadcast message to client:", err)
			conn.Close()
			delete(clients, conn)
		}
	}
}