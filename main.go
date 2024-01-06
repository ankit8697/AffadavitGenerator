package main

import (
	"net/http"
	"log"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/affadavit/s6", buildS6Affadavit)
	mux.HandleFunc("/", renderForm)

	port := os.Getenv("PORT")
	if port == "" {
		port = "0.0.0.0:8000"
	}

	log.Printf("Listening on %s...", port)
	http.ListenAndServe(port, mux)
}