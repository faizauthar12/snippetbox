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
    // SQL Statement
    stmt := `INSERT INTO snippets (title, content, created, expires)
    VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

    // use Exec to insert to the DB
    result, err := m.DB.Exec(stmt, title, content, expires)
    if err != nil {
        return 0, err
    } 

    // Use the LastInsertId() method on the result to get the ID of our
    // newly inserted record in the snippets table.
    id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

    // The ID returned has the type int64, so we convert it to an int type
    // before returning
    return int(id), nil
}

// This will return a specific snippet based on tis id
func (m *SnippetModel) Get(id int) ([]*Snippet, error) {
    return nil, nil
}

// This wil lreturn the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]*Snippet, error) {
    return nil, nil
}
