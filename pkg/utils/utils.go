package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	// Read the body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		return
	}

	// Unmarshal the body
	if err := json.Unmarshal(body, x); err != nil {
		log.Printf("Error unmarshalling body: %v", err)
		return
	}
}
