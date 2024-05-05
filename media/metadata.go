package media

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mucutHTMX/ws"
	"os/exec"
	"strings"
)

type VideoHeatmap struct {
	EndTime   float32 `json:"end_time"`
	StartTime float32 `json:"start_time"`
	Value     float64 `json:"value"` // Updated type to float64 to handle decimal values
}

type UserInput struct {
	From     float32 `json:"from,omitempty"` // Use pointers to make fields optional
	To       float32 `json:"to,omitempty"`
	Duration float32 `json:"duration,omitempty"`
}

type VideoMetadata struct {
	Title          string         `json:"title"`
	Duration       float32        `json:"duration"`        // Assuming duration is in seconds
	DurationString string         `json:"duration_string"` // This might need to be calculated separately if not provided directly
	Heatmap        []VideoHeatmap `json:"heatmap"`
	OriginalUrl    string         `json:"original_url"`
}

type CombinedData struct {
	VideoMetadata VideoMetadata `json:"fetched"`
	UserInput     UserInput     `json:"userInput"`
}

func YtMetadataFromText(uid string, titles []string) {
	for _, title := range titles {
		videoMeta, err := titleToYtMetadata(title)
		if err != nil {
			log.Printf("Error fetching metadata for title '%s': %v", title, err)
			continue
		}

		sendMetadataToWS(uid, videoMeta, title)
	}
}

func sendMetadataToWS(uid string, metadata *VideoMetadata, title string) {
	metaJSON, err := json.Marshal(metadata)

	message := ws.WSMessage{
		Action: "parseTxt",
		Data:   string(metaJSON),
	}

	messageJSON, err := json.Marshal(message)
	if err != nil {
		log.Printf("Error marshalling message to JSON for title '%s': %v", title, err)
	}

	ws.BroadcastMessage(uid, string(messageJSON))
}

func titleToYtMetadata(title string) (*VideoMetadata, error) {
	cmdArgs := []string{
		"--default-search", "ytsearch1:", // Limit to the first search result
		"--dump-json",   // Get the output in JSON format
		"--no-playlist", // Ensure only single video info is returned
		title,           // The search query
	}

	cmd := exec.Command("yt-dlp", cmdArgs...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error executing yt-dlp for title '%s': %v", title, err)
	}

	output = []byte(cleanYtDlpOutput(string(output)))

	var videoMeta VideoMetadata
	err = json.Unmarshal(output, &videoMeta)
	if err != nil {
		log.Printf("Error unmarshaling JSON for title '%s': %v", title, err)
	}

	return &videoMeta, nil
}

func cleanYtDlpOutput(input string) string {
	reader := bufio.NewReader(bytes.NewReader([]byte(input)))

	var cleanData bytes.Buffer
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading output:", err)
				return ""
			}
			break
		}

		if !strings.HasPrefix(line, "WARNING: ") {
			cleanData.WriteString(line)
		}
	}

	return cleanData.String()
}
