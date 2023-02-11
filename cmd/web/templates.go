package main

import (
	"html/template"
	"path/filepath"

	"snippetbox.faizauthar12.github.io/internal/models"
)

type templateData struct {
    CurrentYear int
    Snippet *models.Snippet
    Snippets []*models.Snippet
}

func newTemplateCache() (map[string]*template.Template, error) {
    // Initialize a new map to act as the cache.
    cache := map[string]*template.Template{}

    pages, err := filepath.Glob("./ui/html/pages/*.tmpl")
    if err != nil {
        return nil, err
    }
    // Loop through the page filepaths one-by-one.
    for _, page := range pages {
        
        name := filepath.Base(page)

        ts, err := template.ParseFiles("./ui/html/base.tmpl")
        if err != nil {
            return nil, err
        }

        ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl")
        if err != nil {
            return nil, err
        }

        ts, err = ts.ParseFiles(page)
        if err != nil {
            return nil, err
        }
        
        cache[name] = ts
    }
    // Return the map.
    return cache, nil
}
