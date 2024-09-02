package api

import (
	"AhmadAbdelrazik/mark2right/internal/dictionary"
	"encoding/json"
	"log"
	"net/http"
	"regexp"
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
	InfoLog    *log.Logger
	ErrorLog   *log.Logger
	Dictionary *dictionary.Dictionary
	Regex      []*regexp.Regexp
}

func NewApplication(infoLog, errorLog *log.Logger) (*Application, error) {
	dictionary, err := dictionary.NewDictionary()
	if err != nil {
		return nil, err
	}

	regex, err := CompileRegex()
	if err != nil {
		return nil, err
	}

	app := &Application{
		Dictionary: dictionary,
		Regex:      regex,
		InfoLog:    infoLog,
		ErrorLog:   errorLog,
	}

	return app, nil
}
