package main

import "net/http"

func (app *application) Routes() http.Handler {
	mux := http.NewServeMux()

	return mux
}
