package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// This is the version of the application
const version = "1.0.0"

// This is the config struct that will be used to store the configuration
type config struct {
	port int
	env  string
}

// This is the application struct that will be used to store the configuration and logger and dependencies
type application struct {
	config config
	logger *log.Logger
}

func main() {

	// This is used to parse the command line flags
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.env, "environment", "development", "Environment: {development | staging | production}")

	flag.Parse()

	// This is used to create a new logger instance
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// This is used to create a new application instance
	app := &application{
		config: cfg,
		logger: logger,
	}

	// This is used to create a new http server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// This is used to log the server start
	logger.Printf("Starting %s server on %s", cfg.env, srv.Addr)

	// This is used to start the server
	err := srv.ListenAndServe()

	// This is used to log any errors that occur when starting the server
	logger.Fatal(err)

}
