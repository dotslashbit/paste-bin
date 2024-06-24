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
		Title     string `json:"title"`
		Content   string `json:"content"`
		ExpiresAt string `json:"expires_at"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	expiresAtTime, err := app.ParseExpiresAt(input.ExpiresAt)

	if err != nil {
		app.failedValidationResponse(w, r, map[string]string{"expires_at": "must be in the format YYYY-MM-DD"})
		return
	}

	snippet := &data.Snippet{
		Title:     input.Title,
		Content:   input.Content,
		ExpiresAt: expiresAtTime.Format(time.RFC3339),
	}

	v := validator.New()
	if data.ValidateSnippet(v, snippet); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Snippets.Insert(snippet)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/snippets/%d", snippet.Id))
	err = app.writeJSON(w, http.StatusCreated, envelope{"snippet": snippet}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	// fmt.Fprintf(w, "%+v\n", input)
}

// This is used to display a specific snippet based on the ID (string)
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	snippet, err := app.models.Snippets.Get(int(id))
	if err != nil {
		switch err {
		case data.ErrRecordNotFound:
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"snippet": snippet}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) updateSnippetHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	snippet, err := app.models.Snippets.Get(int(id))
	if err != nil {
		switch err {
		case data.ErrRecordNotFound:
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		Title     string `json:"title"`
		Content   string `json:"content"`
		ExpiresAt string `json:"expires_at"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	expiresAtTime, err := app.ParseExpiresAt(input.ExpiresAt)
	if err != nil {
		app.failedValidationResponse(w, r, map[string]string{"expires_at": "must be in the format YYYY-MM-DD"})
		return
	}

	snippet.Title = input.Title
	snippet.Content = input.Content
	snippet.ExpiresAt = expiresAtTime.Format(time.RFC3339)

	v := validator.New()
	if data.ValidateSnippet(v, snippet); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Snippets.Update(snippet)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"snippet": snippet}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteSnippetHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Snippets.Delete(int(id))
	if err != nil {
		switch err {
		case data.ErrRecordNotFound:
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "snippet deleted successfully"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
