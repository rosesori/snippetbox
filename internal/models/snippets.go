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
	stmt := `INSERT INTO snippets (title, content, created, expires) 
	VALUES (?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	// Use the Exec() method on the embedded connection pool statement.
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	// LastInsertId() on result will get the ID of our newly inserted record
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// The ID returned has the type int64, so we convert it to an int type
	return int(id), nil
}

// Get a specific snippet based on its ID
func (m *SnippetModel) Get(id int) (Snippet, error) {
	return Snippet{}, nil
}

// Latest 10 most recently created snippets will be returned
func (m *SnippetModel) Latest() ([]Snippet, error) {
	return nil, nil
}
