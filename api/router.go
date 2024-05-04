package api

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"net/http"
)

func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/ws/connect", WSHandler)
	router.HandleFunc("/download-siq", DownloadSiqHandler)

	return router
}

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true }, // Allow connections from any origin
}
