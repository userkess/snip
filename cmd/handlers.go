package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from snip")
}

func viewSnip(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Display snip id ", id)
}

func createSnip(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "create a new snip")
}

func createSnipPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Save a new Snip")
}
