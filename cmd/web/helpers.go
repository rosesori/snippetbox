package main

import (
	"net/http"
	"runtime/debug"
)

/*
The serverError helper writes a log entry at Error level
(including the request method and URI as attributes), then sends
a generic 500 Internal Server Error response to the user.
*/
func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
		trace  = string(debug.Stack()) // Get the stack trace, and convert it to a string
	)

	// Log the details of the error message
	app.logger.Error(err.Error(), "method", method, "uri", uri, "trace", trace)
	// Use http.Error() function to send an Internal Server Error response to the user
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// The clientError helper sends a specific status code and corresponding description to the user.
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}
