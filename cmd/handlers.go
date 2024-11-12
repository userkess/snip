package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	files := []string{
		"./html/base.tmpl",
		"./html/nav.tmpl",
		"./html/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Sever Error", http.StatusInternalServerError)
	}
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
