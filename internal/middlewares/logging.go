package middlewares

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/gthomas08/realworld-huma/pkg/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RequestLoggerMiddleware(logger *logger.Logger) echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:      true,
		LogStatus:   true,
		LogLatency:  true,
		LogError:    true,
		HandleError: true,
		BeforeNextFunc: func(c echo.Context) {
			logger.Info("starting request", "endpoint", c.Request().URL.String())
		},
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("request completed", "code", v.Status, "duration", float64(v.Latency)/1e6, "endpoint", v.URI)
			return nil
		},
	})
}

func RecoverMiddleware(logger *logger.Logger, api huma.API, ctx huma.Context, next func(huma.Context)) {
	defer func() {
		if rec := recover(); rec != nil {
			huma.WriteErr(api, ctx, http.StatusInternalServerError, "internal server error")
			return
		}
	}()
	next(ctx)
}
