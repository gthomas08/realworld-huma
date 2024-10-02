package main

import (
	"context"

	"github.com/gthomas08/realworld-huma/config"
	"github.com/gthomas08/realworld-huma/internal/app"
	"github.com/gthomas08/realworld-huma/internal/db/postgres"
	"github.com/gthomas08/realworld-huma/pkg/logger"
)

// main is the entrypoint for the application.
func main() {
	// Create a logger.
	appLogger := logger.NewLogger()

	// Load the application configuration from the specified directory.
	cfg, err := config.LoadConfig("config")
	if err != nil {
		appLogger.Panic("Failed to load the configuration", "error", err)
	}

	// Initialize the PostgreSQL database
	db, err := postgres.NewDB(cfg.Database)
	if err != nil {
		appLogger.Error("Failed to initialize the database", "error", err)
	}
	defer db.Close()

	// Context
	ctx := context.Background()

	// Check if the connection is valid
	if err := db.Conn.PingContext(ctx); err != nil {
		appLogger.Error("Failed to connect to the database", "error", err)
		return
	}

	appLogger.Info("Connected to the database")

	apiApp := app.NewApp(cfg, appLogger, db)

	apiApp.Run()
}
