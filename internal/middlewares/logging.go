package middlewares

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/danielgtaylor/huma/v2"
	"github.com/gthomas08/realworld-huma/pkg/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RequestLoggerMiddleware(logger *logger.Logger) echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURIPath: true,
		LogMethod:  true,
		LogStatus:  true,
		LogLatency: true,
		LogError:   true,
		BeforeNextFunc: func(c echo.Context) {
			logger.Info("starting request", "endpoint", c.Request().URL.String(), "method", c.Request().Method)
		},
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("request completed", "duration", float64(v.Latency)/1e6, "endpoint", v.URIPath, "method", v.Method, "status", v.Status)
			return nil
		},
	})
}

func RecoverMiddleware(logger *logger.Logger, api huma.API, ctx huma.Context, next func(huma.Context)) {
	defer func() {
		if rec := recover(); rec != nil {
			logger.Error("request failed", "error", rec, "endpoint", ctx.Operation().Path, "method", ctx.Operation().Method)
			log.Print(string(debug.Stack()))
			huma.WriteErr(api, ctx, http.StatusInternalServerError, "internal server error")
			return
		}
	}()
	next(ctx)
}
