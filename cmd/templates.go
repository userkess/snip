package main

import (
	"html/template"
	"path/filepath"
	"snip/internal/models"
	"time"
)

type templateData struct {
	CurrentYear int
	Snip        models.Snip
	Snips       []models.Snip
}

func humanDate(t time.Time) string {
	return t.Format("Jan 02 2006 at 03:04 PM")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	// initialize cache map
	cache := map[string]*template.Template{}
	// loop over all tempates
	pages, err := filepath.Glob("./html/*.tmpl")
	if err != nil {
		return nil, err
	}
	//create slice of template filepaths
	for _, page := range pages {
		name := filepath.Base(page)
		// parse base template
		ts, err := template.New(name).Funcs(functions).ParseFiles("./html/base.tmpl")
		if err != nil {
			return nil, err
		}
		ts, err = ts.ParseGlob("./html/*.tmpl")
		if err != nil {
			return nil, err
		}
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}
		// add the template set to cache
		cache[name] = ts
	}
	return cache, nil
}
