package main

import (
	"log"
	"net/http"
)

// Display the home page
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox!"))
}

// Add a snippetView handler function.
func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

// Add a snipperCreate handler function
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func main() {
	/* Use the http.NewServeMux() function to initialize a new servemux
	A servemux (aka a router) stores a mapping between URL routing patterns
	for your application and the corresponding handlers. Usually you have 1
	servemux for your application containing all your routes. */
	mux := http.NewServeMux()

	/* Register the handler functions and corresponding route patterns with
	the servemux */
	mux.HandleFunc("/{$}", home) // Restrict this route to exact matches on / only.
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

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
