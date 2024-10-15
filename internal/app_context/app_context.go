package app_context

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

func SetUserContext(ctx huma.Context, user *jwtkit.UserClaim) huma.Context {
	return huma.WithValue(ctx, UserContextKey, user)
}

func GetUserContext(ctx context.Context) (*jwtkit.UserClaim, error) {
	val := ctx.Value(UserContextKey)
	return jwtkit.ParseUserClaim(val)
}

func SetJWTContext(ctx huma.Context, token string) huma.Context {
	return huma.WithValue(ctx, JWTContextKey, token)
}

func GetJWTContext(ctx context.Context) (string, error) {
	val := ctx.Value(JWTContextKey)
	return val.(string), nil
}
