package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/mustthink/simple-primary/internal/handlers"
)

func readFlags() string {
	var a = flag.String("url", "localhost:8080", "Server url")
	flag.Parse()
	return *a
}

func main() {
	addr := readFlags()
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := handlers.NewApplication(errorLog, addr)

	srv := &http.Server{
		Addr:     addr,
		ErrorLog: errorLog,
		Handler:  app.Routes(),
	}

	log.Println("Starting Hosting on ", addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
