package server

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

type PingResponse struct {
	Body struct {
		Message string `json:"message" example:"Pong!" doc:"Ping message"`
	}
}

func (a *App) routes() *echo.Echo {
	router := echo.New()
	// router.Use(
	// 	middleware.Recoverer, // Handles panics
	// 	middleware.RequestID, // Adds a unique request ID
	// 	middleware.Logger,    // Logs each request
	// 	// Add more global middleware here
	// )

	api := humaecho.New(router, huma.DefaultConfig("My API", "1.0.0"))

	pingRepo := pingRepository.NewPingRepository(a.db)
	pingUc := pingUsecase.NewPingUsecase(pingRepo)
	pingHandler := pingHttp.NewPingHandler(a.logger, pingUc)

	userRepo := userRepository.NewRepository(a.db)
	userUc := userUsecase.NewUsecase(userRepo)
	userHandler := userHttp.NewHandler(a.logger, userUc)

	pingHandler.RegisterPingRoutes(api)
	userHandler.RegisterRoutes(api)

	return router
}
