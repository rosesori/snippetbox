package main

import (
	"fmt"
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

	fmt.Fprintf(w, "Display a specific snipper with ID %d...", id)
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
