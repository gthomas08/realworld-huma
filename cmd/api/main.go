package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gthomas08/realworld-huma/internal/db/sqlite"
	"github.com/gthomas08/realworld-huma/internal/server"
	"github.com/gthomas08/realworld-huma/pkg/logger"
)

type DB struct {
	Conn *sql.DB
}

// main is the entrypoint for the application.
func main() {
	// Create a logger.
	appLogger := logger.NewLogger()

	// // Initialize the SQLite database
	// db, err := sqlite.NewSQLiteDB("dadw")
	// if err != nil {
	// 	appLogger.Error("Failed to initialize the database", "error", err)
	// }
	// defer db.Close()

	// Open an in-memory SQLite database
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	test := sqlite.DB{Conn: db}

	// Context
	ctx := context.Background()

	// Create a table directly
	createTableSQL := `
		CREATE TABLE pings (
		id INTEGER PRIMARY KEY,
		message TEXT NOT NULL
	);`

	// Execute the SQL statement to create the table
	if _, err := db.ExecContext(ctx, createTableSQL); err != nil {
		fmt.Println("Error creating table:", err)
		return
	}

	fmt.Println("Table created successfully!")

	s := server.NewServer(appLogger, &test)

	s.Run()
}
