package middlewares

import (
	"net/http"
	"strings"

	"github.com/danielgtaylor/huma/v2"
	"github.com/gthomas08/realworld-huma/pkg/jwtkit"
)

// AuthMiddleware creates a middleware that will authorize requests based on
// the required scopes for the operation.
func AuthMiddleware(api huma.API, secretKey string) func(ctx huma.Context, next func(huma.Context)) {
	return func(ctx huma.Context, next func(huma.Context)) {
		isAuthorizationRequired := false
		for _, opScheme := range ctx.Operation().Security {
			var ok bool
			if _, ok = opScheme["bearer"]; ok {
				isAuthorizationRequired = true
				break
			}
		}

		if !isAuthorizationRequired {
			next(ctx)
			return
		}

		token := strings.TrimPrefix(ctx.Header("Authorization"), "Bearer ")
		if len(token) == 0 {
			huma.WriteErr(api, ctx, http.StatusUnauthorized, "unauthorized")
			return
		}

		_, err := jwtkit.ValidateToken(token, secretKey)
		if err != nil {
			huma.WriteErr(api, ctx, http.StatusUnauthorized, "unauthorized")
			return
		}

		next(ctx)
	}
}
