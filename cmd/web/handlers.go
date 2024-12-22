package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Display the home page
func home(w http.ResponseWriter, r *http.Request) {
	// Add a "Server: Go" header to the response
	w.Header().Add("Server", "Go")

	/* Initialize a slice containing the paths to the two files. It's important
	to note that the file containing our base template must be the *first* file
	in the slice */
	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}

	/* Use the template.ParseFiles() function to read the files and store the
	templates in a template set. Notice that we use ... to pass the contents of
	the files slice as  variadic arguments */
	ts, err := template.ParseFiles(files...)
	if err != nil {
		// Log the details error message
		log.Print(err.Error())
		/* Use http.Error() function to send an Internal Server Error response
		to the user */
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		// Return from the handler so no subsequent code is executed.
		return
	}
	/* Use the ExecuteTemplate() method to write the content of the "base"
	template as the response body */
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
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
