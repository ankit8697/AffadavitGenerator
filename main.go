package main

import (
	"net/http"
	"log"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/affadavit/s6", buildS6Affadavit)
	mux.HandleFunc("/affadavit/s7", buildS7Affadavit)
	mux.HandleFunc("/", renderForm)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	host := "0.0.0.0:"+port
	log.Printf("Listening on %s...", host)
	log.Fatal(http.ListenAndServe(host, mux))
}