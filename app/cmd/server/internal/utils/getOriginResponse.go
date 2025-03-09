package utils

import (
	"io"
	"log"
	"net/http"
	"os"
)

func GetOriginResponse(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("error with get request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error with covert body response: %v", err)
	}

	if err := os.WriteFile("templates/index.html", body, 0644); err != nil {
		log.Fatalf("error with reading response: %v", err)
	}

	return string(body)
}
