package main

import (
	"errors"
	"fmt"
	"net/http"
	"snip/internal/models"
	"strconv"
	"strings"
	"unicode/utf8"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	snips, err := app.snips.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := app.newTemplateData(r)
	data.Snips = snips

	app.render(w, r, http.StatusOK, "home.tmpl", data)
}

func (app *application) viewSnip(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	snip, err := app.snips.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}
	data := app.newTemplateData(r)
	data.Snip = snip
	app.render(w, r, http.StatusOK, "view.tmpl", data)
}

func (app *application) createSnip(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	app.render(w, r, http.StatusOK, "create.tmpl", data)
}

func (app *application) createSnipPost(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	expires, err := strconv.Atoi(r.PostForm.Get("expires"))
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// check for errors
	checkErrors := make(map[string]string)
	if strings.TrimSpace(title) == "" {
		checkErrors["title"] = "Required field"
	} else if utf8.RuneCountInString(title) > 100 {
		checkErrors["title"] = "Exceeds maxium length"
	}
	if strings.TrimSpace(content) == "" {
		checkErrors["content"] = "Required field"
	}
	if expires != 1 && expires != 7 && expires != 365 {
		checkErrors["expires"] = "This field must equal 1, 7 or 365 days"
	}
	if len(checkErrors) > 0 {
		fmt.Fprint(w, checkErrors)
		return
	}

	// save the snip if no errors
	id, err := app.snips.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/view/%d", id), http.StatusSeeOther)
}
