package api

import (
	"fmt"
	"log"
	"net/http"
)

func WSHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	uid := r.URL.Query().Get("token")
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

func DownloadSiqHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	if token == "" {
		fmt.Printf("No token found in request")
	}
	fmt.Printf(token)

}
