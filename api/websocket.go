package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type WSMessage struct {
	Action string `json:"action"`
	Data   string `json:"data"`
}

var connections = make(map[string]*websocket.Conn)

func WSHandler(w http.ResponseWriter, r *http.Request) {
	uid := r.URL.Query().Get("token")

	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer RemoveConnection(uid)
	//defer conn.Close()

	connections[uid] = conn
	fmt.Println("Connected to websocket with token:", uid)

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("Received: %s", msg)

		// Parse the incoming message into the WSMessage struct
		var wsMsg WSMessage
		err = json.Unmarshal(msg, &wsMsg)
		if err != nil {
			log.Println("Error parsing message:", err)
			continue
		}

		log.Printf(wsMsg.Action)
		// Handle the incoming message
		HandleIncomingMessage(wsMsg, uid)
	}
}

func HandleIncomingMessage(msg WSMessage, uid string) {
	switch msg.Action {
	case "broadcast":
		BroadcastMessage(uid, msg.Data)
	}
}

func RemoveConnection(uid string) {
	delete(connections, uid)
}

func BroadcastMessage(uid string, message string) {
	if err := connections[uid].WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		log.Printf("Error broadcasting message to a connection: %v", err)
	}
}
