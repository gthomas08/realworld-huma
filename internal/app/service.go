package app

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humaecho"
	"github.com/labstack/echo/v4"

	"github.com/gthomas08/realworld-huma/internal/middlewares"
	"github.com/gthomas08/realworld-huma/internal/utils/security"
	"github.com/gthomas08/realworld-huma/pkg/errs"

	userHTTP "github.com/gthomas08/realworld-huma/internal/domain/user/delivery/http"
	userRepository "github.com/gthomas08/realworld-huma/internal/domain/user/repository"
	userUsecase "github.com/gthomas08/realworld-huma/internal/domain/user/usecase"

	profileHTTP "github.com/gthomas08/realworld-huma/internal/domain/profile/delivery/http"
	profileRepository "github.com/gthomas08/realworld-huma/internal/domain/profile/repository"
	profileUsecase "github.com/gthomas08/realworld-huma/internal/domain/profile/usecase"

	articleHTTP "github.com/gthomas08/realworld-huma/internal/domain/article/delivery/http"
	articleRepository "github.com/gthomas08/realworld-huma/internal/domain/article/repository"
	articleUsecase "github.com/gthomas08/realworld-huma/internal/domain/article/usecase"
)

func (app *App) routes() *echo.Echo {
	router := echo.New()
	apiGroup := router.Group("/api")

	config := huma.DefaultConfig(app.cfg.App.Name, app.cfg.App.Version)
	config.Servers = []*huma.Server{
		{URL: "/api"},
	}
	// config.DocsPath = "" // Disable default OpenAPI docs setup
	config.Components.SecuritySchemes = security.GetSecuritySchemes()

	api := humaecho.NewWithGroup(router, apiGroup, config)

	api.UseMiddleware(
		middlewares.LoggerMiddleware(app.logger),
		middlewares.RecoverMiddleware(api, app.logger),
		middlewares.AuthMiddleware(api, app.cfg.JWT.Key),
	)
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

	userRepo := userRepository.NewRepository(app.db)
	userUc := userUsecase.NewUsecase(app.cfg, app.logger, userRepo)
	userHandler := userHTTP.NewHandler(app.cfg, app.logger, userUc)

	profileRepo := profileRepository.NewRepository(app.db)
	profileUc := profileUsecase.NewUsecase(app.cfg, app.logger, profileRepo, userRepo)
	profileHandler := profileHTTP.NewHandler(app.cfg, app.logger, profileUc)

	articleRepo := articleRepository.NewRepository(app.db)
	articleUc := articleUsecase.NewUsecase(app.cfg, app.logger, articleRepo)
	articleHandler := articleHTTP.NewHandler(app.cfg, app.logger, articleUc)

	userHandler.RegisterRoutes(api)
	profileHandler.RegisterRoutes(api)
	articleHandler.RegisterRoutes(api)

	return router
}
