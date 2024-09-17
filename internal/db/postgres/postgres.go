package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// DB wraps the sql.DB struct for Postgres
type DB struct {
	Conn *sql.DB
}

// NewDB initializes a new PostgreSQL database connection using github.com/lib/pq
func NewDB(host, port, user, password, dbname string) (*DB, error) {
	// Format the connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Open the PostgreSQL database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Verify the connection with Ping
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &DB{Conn: db}, nil
}

// Close closes the database connection
func (db *DB) Close() error {
	return db.Conn.Close()
}
