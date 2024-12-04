package websocket

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
	lim  int
	Veri bool
}

type EventHandler func(client *Client, message map[string]interface{})

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	clients       = make(map[*websocket.Conn]Client)
	dataSet       = make(map[string]bool) // Utilisation d'un set pour les donnÃ©es uniques
	lock          = sync.Mutex{}
	eventHandlers = make(map[string]EventHandler)
)

func HandleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error while upgrading connection:", err)
		return
	}
	defer conn.Close()
	//rajouter dans le message le fait de recuperer l'id de l'utilisateur et de le mettre dans le tab
	client := &Client{conn: conn, lim: 10}
	var initialMsg map[string]interface{}
	if err := conn.ReadJSON(&initialMsg); err != nil {
		fmt.Println("Error reading initial message:", err)
	}
	clients[conn] = *client

	SendInitialData(client)
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			lock.Lock()
			delete(clients, conn)
			lock.Unlock()
			break
		}
		var message map[string]interface{}
		if err := json.Unmarshal(msg, &message); err != nil {
			fmt.Println("Error unmarshalling message:", err)
			continue
		}
		if handler, found := eventHandlers[message["type"].(string)]; found {
			handler(client, message)
		}
	}
}

func SendInitialData(client *Client) {
	lock.Lock()
	defer lock.Unlock()
	dataArray := mapToArray(dataSet)
	msg, err := json.Marshal(map[string]interface{}{
		"type": "initial",
		"data": dataArray,
	})
	if err != nil {
		fmt.Println("Error marshalling initial data:", err)
		return
	}
	err = client.conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		fmt.Println("Error sending initial data:", err)
	}
}
