package api

import (
	"encoding/json"
	"io"
	"log"
	"mucutHTMX/media"
	"net/http"
	"strings"
)

func DownloadSiqHandler(w http.ResponseWriter, r *http.Request) {
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

func HandleTxt(w http.ResponseWriter, r *http.Request) {
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
