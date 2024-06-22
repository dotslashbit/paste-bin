package main

import (
	"fmt"
	"net/http"
)

// This is used to create a new snippet
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create a new snippet...")
}

// This is used to display a specific snippet based on the ID (string)
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id == "" {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %s...", id)
}
