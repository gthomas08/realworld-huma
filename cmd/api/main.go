package main

import (
	"log"

	"github.com/gthomas08/realworld-huma/internal/db/sqlite"
	"github.com/gthomas08/realworld-huma/internal/server"
	"github.com/gthomas08/realworld-huma/pkg/logger"
)

// main is the entrypoint for the application.
func main() {
	// Create a logger.
	appLogger := logger.NewLogger()

	// Initialize the SQLite database
	db, err := sqlite.NewSQLiteDB("./main.db")
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}
	defer db.Close()

	s := server.NewServer(appLogger)

	s.Run()
}
