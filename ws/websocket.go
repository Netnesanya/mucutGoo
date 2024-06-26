package ws

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type WSMessage struct {
	Action string `json:"action"`
	Data   string `json:"data"`
}

var Connections = make(map[string]*websocket.Conn)

func Handler(w http.ResponseWriter, r *http.Request) {
	uid := r.URL.Query().Get("token")

	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	//defer RemoveConnection(uid)
	defer conn.Close()

	Connections[uid] = conn

	log.Println("Connected to websocket with token:", uid)

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
	delete(Connections, uid)
}

func BroadcastMessage(uid string, message string) error {
	if err := Connections[uid].WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		log.Printf("Error broadcasting message to a connection: %v", err)
		return err
	}

	return nil
}

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true }, // Allow connections from any origin
}
