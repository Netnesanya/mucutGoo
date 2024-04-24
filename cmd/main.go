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

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*", "https://netnesanya.github.io"}) // Adjust the port to match your Tauri app's port
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

	fmt.Println("Started")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headersOk, originsOk, methodsOk)(router)))
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
