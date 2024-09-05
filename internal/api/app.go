package api

import (
	"encoding/json"
	"log"
	"net/http"
)

// Writes a json response. Used at the end of a handler
// to send the response.
func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

// Simple API response, Used to send simple messages
type APIResponse struct {
	Message string `json:"message"`
}

type Application struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}
