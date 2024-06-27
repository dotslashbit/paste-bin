package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// This is the routes method that will be called in main.go
func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/snippets/:id", app.showSnippet)
	router.HandlerFunc(http.MethodPost, "/v1/snippets", app.createSnippet)
	router.HandlerFunc(http.MethodPatch, "/v1/snippets/:id", app.updateSnippetHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/snippets/:id", app.deleteSnippetHandler)

	return app.recoverPanic(app.enableCORS(app.rateLimit(router)))
}
