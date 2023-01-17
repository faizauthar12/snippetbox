package models

import (
	"database/sql"
	"time"
)

// Define a Snippet type to hold the data for an individual snippet.
// the struct based on Snippet table in MySQL
type Snippet struct {
    ID int
    Title string
    Content string
    Created time.Time
    Expires time.Time
}

// Define a SnippetModel type which wraps a sql.DB connection pool
type SnippetModel struct {
    DB *sql.DB
}

// This will insert a new snippet into database
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
    return 0, nil
}

// This will return a specific snippet based on tis id
func (m *SnippetModel) Get(id int) ([]*Snippet, error) {
    return nil, nil
}

// This wil lreturn the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]*Snippet, error) {
    return nil, nil
}
