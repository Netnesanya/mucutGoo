package media

import (
	"fmt"
	"log"
	"mucutHTMX/ws"
	"os/exec"
	"path/filepath"
)

func DownloadYtPlaylist(playlistUrl string) error {

	return nil
}

func DownloadAudioSegment(url string, startTime, endTime float32, title string, uid string) error {
	if url == "" {
		return fmt.Errorf("URL is empty, cannot download segment")
	}

	outputPath := filepath.Join(uid, title)

	fmt.Println(url, startTime, endTime, outputPath)

	cmdArgs := []string{
		url,
		"-x",                    // Extract audio
		"--audio-format", "mp3", // Specify audio format, adjust as needed
		"-o", outputPath, // Specify the output path and filename
		"--external-downloader", "ffmpeg",
		"--external-downloader-args", fmt.Sprintf("ffmpeg_i:-ss %.2f -to %.2f", startTime, endTime),
	}

	cmd := exec.Command("yt-dlp", cmdArgs...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to download audio segment: %s, error: %v", string(output), err)
	} else {
		successMessage := fmt.Sprintf("Successfully downloaded audio segment: %s", outputPath)
		if err := ws.BroadcastMessage(uid, successMessage); err != nil {
			log.Printf("Failed to send success message: %v", err)
		}
	}

	return nil
}
