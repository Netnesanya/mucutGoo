package api

import (
	"encoding/json"
	"io"
	"log"
	"mucutHTMX/media"
	"net/http"
	"path/filepath"
	"strings"
)

func DownloadSiqHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	uid := r.Header.Get("Authorization")
	if uid == "" {
		http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
		return
	}

	siqName := "packGovna"
	filePath, err := media.CreateSIQPackage(siqName, uid)
	if err != nil {
		log.Printf("Error creating SIQ package: %v", err)
		http.Error(w, "Error creating SIQ package", http.StatusInternalServerError)
		return
	}

	// Set the headers to prompt the user to download the file
	w.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(filePath))
	w.Header().Set("Content-Type", "application/octet-stream")

	// Serve the file
	http.ServeFile(w, r, filePath)
}

func PrepAudioHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestData []media.CombinedData

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &requestData); err != nil {
		log.Printf("Error unmarshaling request body: %v", err)
		http.Error(w, "Error processing request body", http.StatusBadRequest)
		return
	}

	uid := r.Header.Get("Authorization")

	err = media.DownloadAudioFromList(requestData, uid)
	if err != nil {
		log.Printf("Download failed: %s", err)
	}
}

func TxtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	uid := r.Header.Get("Authorization")
	titles := strings.Split(string(body), "\n")
	media.YtMetadataFromText(uid, titles)

	w.WriteHeader(http.StatusOK)
	return
}
