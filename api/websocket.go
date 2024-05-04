package api

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var connections = make(map[string]*websocket.Conn)

func WSHandler(w http.ResponseWriter, r *http.Request) {
	uid := r.URL.Query().Get("token")

	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer RemoveConnection(uid)
	defer conn.Close()

	connections[uid] = conn
	fmt.Println("Connected to websocket with token:", uid)

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("Received: %s", msg)
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
