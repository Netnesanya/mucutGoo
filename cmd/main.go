package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"log"
	"mucutHTMX/api"
	"net/http"
	"os/exec"
)

func main() {
	checkInstallations()

	router := api.Router()
	const address = ":8080"

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

	fmt.Println("Started")
	log.Fatal(http.ListenAndServe(address, handlers.CORS(headersOk, originsOk, methodsOk)(router)))
}

func checkInstallations() {
	// Check ffmpeg installation
	if err := exec.Command("ffmpeg", "-version").Run(); err != nil {
		log.Println("ffmpeg is not installed")
	} else {
		log.Println("ffmpeg is installed")
	}

	// Check yt-dlp installation
	if err := exec.Command("yt-dlp", "--version").Run(); err != nil {
		log.Println("yt-dlp is not installed")
	} else {
		log.Println("yt-dlp is installed")
	}
}
