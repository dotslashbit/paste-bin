package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// This is used to read the ID parameter from the URL
func (app *application) readIDParam(r *http.Request) (string, error) {
	params := httprouter.ParamsFromContext(r.Context())
	id := params.ByName("id")

	return id, nil
}
