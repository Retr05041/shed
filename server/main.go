package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var DEBUG = true

// CommandRequest will be the initial package sent by a post request
type CommandRequest struct {
	Command string          `json:"command"`
	Payload json.RawMessage `json:"payload"` // It will demarshaled but remain JSON
}

// State?
type Payload struct {
	Nickname string `json:"nickname"`
}

// initiates connection with client, either creates a lobby for the player, or joins them to an already made one
// also creates player info to initiat session
func connectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	// Aquire the commandRequest
	var cmdReq CommandRequest
	err := json.NewDecoder(r.Body).Decode(&cmdReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if DEBUG {
		fmt.Printf("Received command request: '%s' with payload: %s\n", cmdReq.Command, cmdReq.Payload)
	}

	// Validate command exists
	if cmdReq.Command == "" {
		http.Error(w, "Must supply command", http.StatusBadRequest)
		return
	}

	// Identify command - will probably need to split to another file for just commands again
	var responseMessage string
	switch cmdReq.Command {
	case "connect":
		var payload Payload
		if err := json.Unmarshal(cmdReq.Payload, &payload); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if payload.Nickname == "" {
			http.Error(w, "Nickname cannot be empty", http.StatusBadRequest)
			return
		}

		responseMessage = fmt.Sprintf("Hello, %s!", payload.Nickname)
	default:
		http.Error(w, "Invalid command.", http.StatusBadRequest)
		return
	}

	// Send response to just that writer
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(responseMessage))
}

func main() {
	http.HandleFunc("/connect", connectHandler)
	fmt.Println("Server is listening on :8080...")
	http.ListenAndServe(":8080", nil)
}
