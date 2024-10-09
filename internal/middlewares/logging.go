package middlewares

import (
	"log"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/gthomas08/realworld-huma/pkg/logger"
)

// LoggerMiddleware logs information about the request before and after calling the next handler.
//
// It logs the start of the request with the endpoint and method, then calls the next handler.
// After the next handler has finished, it logs the end of the request with the duration, status, endpoint, and method.
//
// The duration is logged in milliseconds as a float64.
func LoggerMiddleware(logger *logger.Logger) func(ctx huma.Context, next func(huma.Context)) {
	return func(ctx huma.Context, next func(huma.Context)) {
		// Log the start of the request
		startTime := time.Now()
		logger.Info("starting request", "endpoint", ctx.Operation().Path, "method", ctx.Operation().Method)

		// Call the next handler
		next(ctx)

		// Log the end of the request and the duration
		duration := time.Since(startTime)
		logger.Info("request completed", "duration", float64(duration)/1e6, "endpoint", ctx.Operation().Path, "method", ctx.Operation().Method, "status", ctx.Status())

	}
}

// RecoverMiddleware recovers from panics and logs the error and stack trace.
//
// This middleware can be used to catch any panics that occur in the handlers and log them.
// After logging the error, it will write a 500 Internal Server Error response to the client.
func RecoverMiddleware(api huma.API, logger *logger.Logger) func(ctx huma.Context, next func(huma.Context)) {
	return func(ctx huma.Context, next func(huma.Context)) {
		defer func() {
			if rec := recover(); rec != nil {
				logger.Error("request failed", "error", rec)
				log.Print(string(debug.Stack()))
				huma.WriteErr(api, ctx, http.StatusInternalServerError, "internal server error")
				return
			}
		}()
		next(ctx)
	}
}
