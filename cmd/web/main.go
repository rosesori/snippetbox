package main

import (
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rosesori/snippetbox/internal/models"
	"log/slog"
	"net/http"
	"os"
)

// Define an application struct to hold the application-wide dependencies
type application struct {
	logger   *slog.Logger
	snippets *models.SnippetModel
}

/*
	 We've limited the responsibilities of the main() function to be:
		- Parsing the runtime configuration settings for the application
		- Establishing the dependencies for the handlers
		- Starting the HTTP server
*/
func main() {
	// Define command line flags
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "web:initialpass@/snippetbox?parseTime=true", "MySQL data source name")

	/* Use the flag.Parse() function to parse the command-line flag.
	This reads in the command line flag values and assign them to corresponding variables .*/
	flag.Parse()

	/* Initialize a new structured logger which writes to the standard out
	stream and uses the default settings. */
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	// Defer so that the connection pool is closed before the main() function exits
	defer db.Close()

	/* Initialize a new instance of the application struct containing the
	dependencies.
	snippets is a models.SnippetModel instance containing the connection pool.
	*/
	app := &application{
		logger:   logger,
		snippets: &models.SnippetModel{DB: db},
	}

	/* The value returned from flag.String() is a pointer to the flag value,
	so we need to dereference the pointer. */
	logger.Info("starting server", "addr", *addr)

	// We use the http.ListenAndServe() function to start a new web server.
	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

// The openDB() function wraps sql.Open() and returns a sql.DB connection pool for a given DSN.
func openDB(dsn string) (*sql.DB, error) {
	// Use sql.Open() to create a connection pool
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Use the Ping() method to check that the connection is working
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
