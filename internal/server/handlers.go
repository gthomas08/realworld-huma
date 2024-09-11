package server

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humaecho"
	"github.com/labstack/echo/v4"

	pingHttp "github.com/gthomas08/realworld-huma/internal/domain/ping/delivery/http"
	pingUsecase "github.com/gthomas08/realworld-huma/internal/domain/ping/usecase"
)

type PingResponse struct {
	Body struct {
		Message string `json:"message" example:"Pong!" doc:"Ping message"`
	}
}

func (s *Server) routes() *echo.Echo {
	router := echo.New()
	// router.Use(
	// 	middleware.Recoverer, // Handles panics
	// 	middleware.RequestID, // Adds a unique request ID
	// 	middleware.Logger,    // Logs each request
	// 	// Add more global middleware here
	// )

	api := humaecho.New(router, huma.DefaultConfig("My API", "1.0.0"))

	pingUc := pingUsecase.NewPingUsecase()
	pingHandler := pingHttp.NewPingHandler(s.logger, pingUc)

	pingHandler.RegisterPingRoutes(api)

	return router
}
