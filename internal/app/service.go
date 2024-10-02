package app

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humaecho"
	"github.com/labstack/echo/v4"

	pingHttp "github.com/gthomas08/realworld-huma/internal/domain/ping/delivery/http"
	pingRepository "github.com/gthomas08/realworld-huma/internal/domain/ping/repository"
	pingUsecase "github.com/gthomas08/realworld-huma/internal/domain/ping/usecase"

	userHttp "github.com/gthomas08/realworld-huma/internal/domain/user/delivery/http"
	userRepository "github.com/gthomas08/realworld-huma/internal/domain/user/repository"
	userUsecase "github.com/gthomas08/realworld-huma/internal/domain/user/usecase"
)

func (app *App) routes() *echo.Echo {
	router := echo.New()
	// router.Use(
	// 	middleware.Recoverer, // Handles panics
	// 	middleware.RequestID, // Adds a unique request ID
	// 	middleware.Logger,    // Logs each request
	// 	// Add more global middleware here
	// )

	api := humaecho.New(router, huma.DefaultConfig(app.cfg.App.Name, app.cfg.App.Version))

	pingRepo := pingRepository.NewPingRepository(app.db)
	pingUc := pingUsecase.NewPingUsecase(pingRepo)
	pingHandler := pingHttp.NewPingHandler(app.logger, pingUc)

	userRepo := userRepository.NewRepository(app.db)
	userUc := userUsecase.NewUsecase(app.logger, userRepo)
	userHandler := userHttp.NewHandler(app.logger, userUc)

	pingHandler.RegisterPingRoutes(api)
	userHandler.RegisterRoutes(api)

	return router
}
