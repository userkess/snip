package main

import (
	"errors"
	"fmt"
	"net/http"
	"snip/internal/models"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	snips, err := app.snips.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	app.render(w, r, http.StatusOK, "home.tmpl", templateData{
		Snips: snips,
	})
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
	app.render(w, r, http.StatusOK, "view.tmpl", templateData{
		Snip: snip,
	})
}

func (app *application) createSnip(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "create a new snip")
}

func (app *application) createSnipPost(w http.ResponseWriter, r *http.Request) {
	title := "0 snail"
	content := "0 snail\nClimb Mount Fuji,\n,But slowly, slowly!\n\n- Kobayashi Issa"
	expires := 7
	id, err := app.snips.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/view/%d", id), http.StatusSeeOther)
}
