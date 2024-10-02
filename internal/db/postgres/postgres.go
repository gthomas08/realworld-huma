package postgres

import (
	"database/sql"
	"fmt"

	"github.com/gthomas08/realworld-huma/config"
	_ "github.com/lib/pq"
)

// DB wraps the sql.DB struct for Postgres
type DB struct {
	Conn *sql.DB
}

// NewDB initializes a new PostgreSQL database connection.
func NewDB(dbConfig config.DatabaseConfig) (*DB, error) {
	// Format the connection string
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Name)

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
