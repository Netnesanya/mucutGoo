package api

import (
	"github.com/gorilla/mux"
	"mucutHTMX/ws"
)

func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/ws/connect", ws.Handler)
	router.HandleFunc("/siq", DownloadSiqHandler)
	router.HandleFunc("/parse-txt", HandleTxt).Methods("POST")

	return router
}
