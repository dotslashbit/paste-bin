package main

import (
	"fmt"
	"net/http"
)

// This is used to check the health of the application
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Status: OK")
	fmt.Fprintf(w, "Environment: %s\n", app.config.env)
	fmt.Fprintf(w, "Version: %s\n", version)
}
