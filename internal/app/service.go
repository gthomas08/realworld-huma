package app

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humaecho"
	"github.com/labstack/echo/v4"

	"github.com/gthomas08/realworld-huma/internal/middlewares"
	"github.com/gthomas08/realworld-huma/pkg/errs"

	pingHttp "github.com/gthomas08/realworld-huma/internal/domain/ping/delivery/http"
	pingRepository "github.com/gthomas08/realworld-huma/internal/domain/ping/repository"
	pingUsecase "github.com/gthomas08/realworld-huma/internal/domain/ping/usecase"

	userHttp "github.com/gthomas08/realworld-huma/internal/domain/user/delivery/http"
	userRepository "github.com/gthomas08/realworld-huma/internal/domain/user/repository"
	userUsecase "github.com/gthomas08/realworld-huma/internal/domain/user/usecase"
)

func (app *App) routes() *echo.Echo {
	router := echo.New()

	api := humaecho.New(router, huma.DefaultConfig(app.cfg.App.Name, app.cfg.App.Version))
	api.UseMiddleware(
		middlewares.LoggerMiddleware(app.logger),
		middlewares.RecoverMiddleware(api, app.logger),
	)
	huma.NewError = errs.NewError

	pingRepo := pingRepository.NewPingRepository(app.db)
	pingUc := pingUsecase.NewPingUsecase(pingRepo)
	pingHandler := pingHttp.NewPingHandler(app.logger, pingUc)

	userRepo := userRepository.NewRepository(app.db)
	userUc := userUsecase.NewUsecase(app.cfg, app.logger, userRepo)
	userHandler := userHttp.NewHandler(app.cfg, app.logger, userUc)

	pingHandler.RegisterPingRoutes(api)
	userHandler.RegisterRoutes(api)

	return router
}
