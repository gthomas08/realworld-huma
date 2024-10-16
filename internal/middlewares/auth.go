package middlewares

import (
	"net/http"
	"strings"

	"github.com/danielgtaylor/huma/v2"
	"github.com/gthomas08/realworld-huma/internal/ctxkit"
	"github.com/gthomas08/realworld-huma/pkg/jwtkit"
)

// AuthMiddleware creates a middleware that will authorize requests based on
// the required scopes for the operation.
func AuthMiddleware(api huma.API, secretKey string) func(ctx huma.Context, next func(huma.Context)) {
	return func(ctx huma.Context, next func(huma.Context)) {
		// Check if any of the operation's schemes require authorization
		isAuthorizationRequired := false
		for _, opScheme := range ctx.Operation().Security {
			var ok bool
			if _, ok = opScheme["bearer"]; ok {
				isAuthorizationRequired = true
				break
			}
		}

		// If authorization is not required, call the next middleware function
		if !isAuthorizationRequired {
			next(ctx)
			return
		}

		// Get the JWT from the Authorization header
		token := strings.TrimPrefix(ctx.Header("Authorization"), "Bearer ")
		if len(token) == 0 {
			huma.WriteErr(api, ctx, http.StatusUnauthorized, "no token provided")
			return
		}

		// Verify the JWT using the provided secret key
		validatedToken, err := jwtkit.ValidateToken(token, secretKey)
		if err != nil {
			huma.WriteErr(api, ctx, http.StatusUnauthorized, "unauthorized")
			return
		}

		// Get the user claim from the JWT
		user, exists := validatedToken.Get("user")
		if !exists {
			huma.WriteErr(api, ctx, http.StatusUnauthorized, "user claim not found")
			return
		}

		// Parse the user claim into a struct
		userClaim, err := jwtkit.ParseUserClaim(user)
		if err != nil {
			huma.WriteErr(api, ctx, http.StatusUnauthorized, "invalid user claim")
			return
		}

		// Set the user claim and JWT in the context
		ctx = ctxkit.SetUserContext(ctx, userClaim)
		ctx = ctxkit.SetJWTContext(ctx, token)

		// Call the next middleware function
		next(ctx)
	}
}
