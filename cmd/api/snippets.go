package main

import (
	"fmt"
	"net/http"
	"time"

	"dev.dotslashbit.paste-bin/internal/data"
	"dev.dotslashbit.paste-bin/internal/validator"
)

// This is used to create a new snippet
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	snippet := &data.Snippet{
		Title:   input.Title,
		Content: input.Content,
	}

	v := validator.New()
	if data.ValidateSnippet(v, snippet); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

// This is used to display a specific snippet based on the ID (string)
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id == "" {
		app.notFoundResponse(w, r)
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
		app.serverErrorResponse(w, r, err)
	}

}
