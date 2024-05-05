package api

import (
	"fmt"
	"io"
	"mucutHTMX/media"
	"net/http"
	"strings"
)

func DownloadSiqHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	if token == "" {
		fmt.Printf("No token found in request")
	}
	fmt.Printf(token)
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
