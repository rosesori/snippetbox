package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	/* Define a new command line flag with the name "addr", a default
	value of ":4000", and a short description. */
	addr := flag.String("addr", ":4000", "HTTP network address")

	/* Use the flag.Parse() function to parse the command-line flag.
	This reads in the command line flag value and assigns it to the addr variable.*/
	flag.Parse()

	/* Use the http.NewServeMux() function to initialize a new servemux.
	A servemux (aka a router) stores a mapping between URL routing patterns
	for your application and the corresponding handlers. Usually you have 1
	servemux for your application containing all your routes. */
	mux := http.NewServeMux()

	/* Create a file server which serves files out of the "./ui/static" directory.
	Note that the path given to the http.Dir function is relative to the project
	directory root. */
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	/* Use the mux.Handle() function to register the file server as the handler
	for all URL paths that start with "/static/". For matching paths, we strip
	the "/static" prefix before the request reaches the file server */
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	/* Register the handler functions and corresponding route patterns with
	the servemux */
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippert/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	/* Print a log message to say that the server is starting.
	The value returned from flag.String() is a pointer to the flag value,
	so we need to dereference the pointer. */
	log.Printf("starting server on %s", *addr)

	/* Use the http.ListenAndServe() function to start a new web server.
	We pass in two parameters:
	  - the TCP network address to listen on (in this case ":4000")
	  - and the servemux we just created.
	If http.ListenAndServe() returns an error, we use the log.Fatal() function
	to log the error message and exit. Note that any error returned by
	http.ListenAndServe() will always be non-nil. */
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
