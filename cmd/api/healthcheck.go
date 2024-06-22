package main

import (
	"net/http"
	"time"
)

// This is used to check the health of the application
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	env := envelope{
		"status": "available",
		"system_info": map[string]interface{}{
			"environment": app.config.env,
			"server_time": time.Now(),
			"version":     version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, envelope{"environment": env}, nil)
	if err != nil {
		app.logger.Println(err)
	}
}
