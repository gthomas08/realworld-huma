package app

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humaecho"
	"github.com/labstack/echo/v4"

	"github.com/gthomas08/realworld-huma/internal/middlewares"
	"github.com/gthomas08/realworld-huma/internal/utils/security"
	"github.com/gthomas08/realworld-huma/pkg/errs"
)

func (app *App) routes() *echo.Echo {
	router := echo.New()
	apiGroup := router.Group("/api")

	// Set up Huma config
	config := huma.DefaultConfig(app.cfg.App.Name, app.cfg.App.Version)
	config.Servers = []*huma.Server{
		{URL: "/api"},
	}
	// config.DocsPath = "" // Disable default OpenAPI docs setup
	config.Components.SecuritySchemes = security.GetSecuritySchemes()

	api := humaecho.NewWithGroup(router, apiGroup, config)

	// Set up middleware
	api.UseMiddleware(
		middlewares.LoggerMiddleware(app.logger),
		middlewares.RecoverMiddleware(api, app.logger),
		middlewares.AuthMiddleware(api, app.cfg.JWT.Key),
	)

	// Set up custom error handling
	huma.NewError = errs.NewError

	// router.GET("/api/docs", func(c echo.Context) error {
	// 	return c.HTML(http.StatusOK, `<!doctype html>
	// 	<html>
	// 	  <head>
	// 		<title>API Reference</title>
	// 		<meta charset="utf-8" />
	// 		<meta name="viewport" content="width=device-width, initial-scale=1" />
	// 	  </head>
	// 	  <body>
	// 		<script id="api-reference" data-url="/api/openapi.json"></script>
	// 		<script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>
	// 	  </body>
	// 	</html>`)
	// })

	// Register domain-specific handlers
	app.registerRoutes(api)

	return router
}
