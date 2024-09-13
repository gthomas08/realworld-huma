package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// DB wraps the sql.DB struct
type DB struct {
	Conn *sql.DB
}

// NewSQLiteDB initializes a new SQLite database using github.com/mattn/go-sqlite3
func NewSQLiteDB(dbPath string) (*DB, error) {
	// Open the SQLite database
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	return &DB{Conn: db}, nil
}

// Close closes the database connection
func (db *DB) Close() error {
	return db.Conn.Close()
}
