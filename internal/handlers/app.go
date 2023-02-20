package handlers

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

type Application struct {
	errorLog *log.Logger
	url      string
	client   *http.Client
}

func NewApplication(errorLog *log.Logger, url string) *Application {
	return &Application{
		errorLog: errorLog,
		url:      url,
		client:   &http.Client{},
	}
}

func (app *Application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Println(trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *Application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *Application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
