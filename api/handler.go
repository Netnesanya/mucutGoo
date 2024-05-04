package api

import (
	"fmt"
	"net/http"
)

func DownloadSiqHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	if token == "" {
		fmt.Printf("No token found in request")
	}
	fmt.Printf(token)
}
