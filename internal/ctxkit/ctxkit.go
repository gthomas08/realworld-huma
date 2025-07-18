package ctxkit

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/gthomas08/realworld-huma/pkg/jwtkit"
)

type contextKey string

const (
	UserContextKey contextKey = "user"
	JWTContextKey  contextKey = "jwt"
)

// SetUserContext adds the provided UserClaim to the context.
func SetUserContext(ctx huma.Context, user *jwtkit.UserClaim) huma.Context {
	return huma.WithValue(ctx, UserContextKey, user)
}

// GetUserContext retrieves the UserClaim from the provided context.
func GetUserContext(ctx context.Context) (*jwtkit.UserClaim, bool) {
	val := ctx.Value(UserContextKey)
	user, ok := val.(*jwtkit.UserClaim)
	return user, ok
}

// SetJWTContext adds the provided JWT token to the context.
func SetJWTContext(ctx huma.Context, token string) huma.Context {
	return huma.WithValue(ctx, JWTContextKey, token)
}

// GetJWTContext retrieves the JWT token from the provided context.
func GetJWTContext(ctx context.Context) (string, bool) {
	val := ctx.Value(JWTContextKey)
	token, ok := val.(string)
	return token, ok
}
