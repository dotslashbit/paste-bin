package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"dev.dotslashbit.paste-bin/internal/data"
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

	snippet := data.Snippet{
		Id:       id,
		Title:    "An old silent pond",
		Content:  "An old silent pond...",
		ExpireAt: time.Now().Add(5 * time.Minute),
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"snippet": snippet}, nil)
	if err != nil {
		errors.New("error writing JSON response")
	}

}
