package models

import (
	"database/sql"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// SnippetModel type wraps a sql.DB connection pool.
// A database model is the same as a "service layer" or "data access layer"
type SnippetModel struct {
	DB *sql.DB
}

// Insert a new snippet into the database
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	return 0, nil
}

// Get a specific snippet based on its ID
func (m *SnippetModel) Get(id int) (Snippet, error) {
	return Snippet{}, nil
}

// Latest 10 most recently created snippets will be returned
func (m *SnippetModel) Latest() ([]Snippet, error) {
	return nil, nil
}
