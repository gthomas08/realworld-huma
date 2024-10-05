package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/gthomas08/realworld-huma/config"
	"github.com/gthomas08/realworld-huma/internal/db/postgres"
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
		// Set up the HTTP server with the application's routes and sensible timeout settings.
		server := &http.Server{
			Addr:         fmt.Sprintf(":%d", app.cfg.Server.Port),
			Handler:      app.routes(),
			IdleTimeout:  time.Minute,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			ErrorLog:     app.logger.NewLogLogger(),
		}

		// Hook to start the server.
		hooks.OnStart(func() {
			app.logger.Info("starting server", "port", app.cfg.Server.Port, "env", app.cfg.App.Env)
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				app.logger.Error("failed to start server", "error", err.Error())
			}
		})

		// Hook to stop the server.
		hooks.OnStop(func() {
			// Give the server some seconds to gracefully shut down.
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := server.Shutdown(ctx); err != nil {
				app.logger.Error("failed to shutdown server", "error", err.Error())
			}
		})
	})

	// Run the CLI. If no commands are passed, it starts the server.
	cli.Run()
}
