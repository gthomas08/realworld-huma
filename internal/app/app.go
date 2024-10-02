package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/gthomas08/realworld-huma/config"
	"github.com/gthomas08/realworld-huma/internal/db/postgres"
	"github.com/gthomas08/realworld-huma/pkg/errs"
	"github.com/gthomas08/realworld-huma/pkg/logger"
)

// Options contains all the configurable options for the application.
type Options struct{}

type App struct {
	cfg    *config.Config
	logger *logger.Logger
	db     *postgres.DB
}

func NewApp(cfg *config.Config, logger *logger.Logger, db *postgres.DB) *App {
	return &App{cfg: cfg, logger: logger, db: db}
}

func (app *App) Run() {
	// Create a CLI app with the provided options.
	cli := humacli.New(func(hooks humacli.Hooks, options *Options) {
		// Set up the logger.
		logger := logger.NewLogger()

		// Set up the error handler.
		huma.NewError = errs.NewError

		// Set up the HTTP server with the application's routes and sensible timeout settings.
		server := &http.Server{
			Addr:         fmt.Sprintf(":%d", app.cfg.Server.Port),
			Handler:      app.routes(),
			IdleTimeout:  time.Minute,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			ErrorLog:     logger.NewLogLogger(),
		}

		// Hook to start the server.
		hooks.OnStart(func() {
			logger.Info("Starting server", "port", app.cfg.Server.Port, "env", app.cfg.App.Env)
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				logger.Error("Failed to start server", "error", err)
			}
		})

		// Hook to stop the server.
		hooks.OnStop(func() {
			// Give the server some seconds to gracefully shut down.
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := server.Shutdown(ctx); err != nil {
				logger.Error("Failed to shutdown server", "error", err)
			}
		})
	})

	// Run the CLI. If no commands are passed, it starts the server.
	cli.Run()
}
