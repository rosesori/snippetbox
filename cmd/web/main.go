package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

// Define an application struct to hold the application-wide dependencies
type application struct {
	logger *slog.Logger
}

/*
	 We've limited the responsibilities of the main() function to be:
		- Parsing the runtime configuration settings for the application
		- Establishing the dependencies for the handlers
		- Starting the HTTP server
*/
func main() {
	/* Define a new command line flag with the name "addr", a default
	value of ":4000", and a short description. */
	addr := flag.String("addr", ":4000", "HTTP network address")

	/* Use the flag.Parse() function to parse the command-line flag.
	This reads in the command line flag value and assigns it to the addr variable.*/
	flag.Parse()

	/* Initialize a new structured logger which writes to the standard out
	stream and uses the default settings. */
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	/* Initialize a new instance of the application struct containing the
	dependencies */
	app := &application{logger: logger}

	/* The value returned from flag.String() is a pointer to the flag value,
	so we need to dereference the pointer. */
	logger.Info("starting server", "addr", *addr)

	// We use the http.ListenAndServe() function to start a new web server.
	err := http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
