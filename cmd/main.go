package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /view/{id}", viewSnip)
	mux.HandleFunc("GET /create", createSnip)
	mux.HandleFunc("POST /create", createSnipPost)
	log.Print("starting server on port 8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
