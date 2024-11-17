package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	files := []string{
		"./html/base.tmpl",
		"./html/nav.tmpl",
		"./html/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) viewSnip(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Display snip id ", id)
}

func (app *application) createSnip(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "create a new snip")
}

func (app *application) createSnipPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Save a new Snip")
}
