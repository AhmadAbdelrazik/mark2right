package api

import "net/http"

func (a *Application) Routes() http.Handler {
	mux := http.NewServeMux()

	return mux
}
