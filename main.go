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
		port = "8080"
	}

	log.Printf("Listening on %s...", "0.0.0.0:"+port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, mux))
}