package main

import (
	"net/http"
	"log"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/affadavit/s6", buildS6Affadavit)
	mux.HandleFunc("/", renderForm)

	log.Print("Listening...")
	http.ListenAndServe(":8000", mux)
}