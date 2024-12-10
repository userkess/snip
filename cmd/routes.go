package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app application) routes() http.Handler {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /view/{id}", app.viewSnip)
	mux.HandleFunc("GET /create", app.createSnip)
	mux.HandleFunc("POST /create", app.createSnipPost)
	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)
	return standard.Then(mux)
}
