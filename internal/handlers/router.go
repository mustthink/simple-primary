package handlers

import "net/http"

func (app *Application) Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/check", app.checkPrimaries)

	return mux
}
