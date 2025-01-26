package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

// Display the home page
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Add a "Server: Go" header to the response
	w.Header().Add("Server", "Go")

	// Return the latest 10 snippets
	// snippets, err := app.snippets.Latest()
	// if err != nil {
	// 	app.serverError(w, r, err)
	// 	return
	// }

	// for _, snippet := range snippets {
	// 	fmt.Fprintf(w, "%+v\n", snippet)
	// }

	/* Initialize a slice containing the paths to the two files. It's important
	to note that the file containing our base template must be the *first* file
	in the slice */
	files := []string{
		"./ui/html/base.gohtml",
		"./ui/html/partials/nav.gohtml",
		"./ui/html/pages/home.gohtml",
	}

	/* Use the template.ParseFiles() function to read the files and store the
	templates in a template set. Notice that we use ... to pass the contents of
	the files slice as  variadic arguments */
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return // Return from the handler so no subsequent code is executed.
	}

	/* Use the ExecuteTemplate() method to write the content of the "base"
	template as the response body */
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}

// Display a specific snippet
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id wildcard from the request using r.PathValue()
	// Try to convert id to an integer using strconv.Atoi()
	// If it cannot be converted to an integer, or the value is less than 1, return 404
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// Use the SnippetModel.Get() method to retrieve the data for a specific record based on its ID
	// Return 404 Not Found if no matching record is found
	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	// Slice with paths to the view.tmpl file
	files := []stirng{
		"./ui/html/base.gohtml",
		"./ui/html/partials/nav.gohtml",
		"./ui/html/pages/view.gohtml",
	}

	// Parse the template files
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Execute the template set, passing in the snippet data
	err = ts.ExecuteTemplate(w, "base", snippet)
	if err != nil {
		app.serverError(w, r, err)
	}

	// Write the snippet data as a plain-text HTTP response body
	fmt.Fprintf(w, "%+v", snippet)
}

// Display a form for creating a new snippet
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

// Save a new snippet
func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	// Create some variables holding dummy data
	title := "O snail"
	content := "O snail \nClimb Mount Fuji, \nBut slowly, slowly!\n\n- Kobayashi Issa"
	expires := 7

	// Pass the data to the SnippetModel.Insert() method, receiving the ID of the new record back.
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Redirect the user to the relevant page for the snippet.
	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
