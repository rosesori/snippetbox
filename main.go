package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Display the home page
func home(w http.ResponseWriter, r *http.Request) {
	// Add a "Server: Go" header to the response
	w.Header().Add("Server", "Go")

	w.Write([]byte("Hello from Snippetbox!"))
}

// Display a specific snippet
func snippetView(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id wildcard from the request using r.PathValue()
	// Try to convert id to an integer using strconv.Atoi()
	id, err := strconv.Atoi(r.PathValue("id"))

	// If cannot be converted to an integer,
	// or the value is less than 1, return 404
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	msg := fmt.Sprintf("Display a specific snippet with ID %d...", id)
	w.Write([]byte(msg))
}

// Display a form for creating a new snippet
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

// Save a new snippet
func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	// Use the w.WriteHeader() method to send a 201 status code.
	w.WriteHeader(http.StatusCreated)

	// Then w.Write() method to write the response body as normal.
	w.Write([]byte("Save a new snippet..."))
}

func main() {
	/* Use the http.NewServeMux() function to initialize a new servemux
	A servemux (aka a router) stores a mapping between URL routing patterns
	for your application and the corresponding handlers. Usually you have 1
	servemux for your application containing all your routes. */
	mux := http.NewServeMux()

	/* Register the handler functions and corresponding route patterns with
	the servemux */
	mux.HandleFunc("GET /{$}", home) // Restrict this route to exact matches on / only.
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	// Print a log message to say that the server is starting.
	log.Print("Starting server on :4000")

	/* Use the http.ListenAndServe() function to start a new web server.
	We pass in two parameters: the TCP network address to listen on (in this case ":4000")
	and the servemux we just created. If http.ListenAndServe() returns an error we use the log.Fatal()
	function to log the error message and exit. Note that any error returned by http.ListenAndServe()
	will always be non-nil. */
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
