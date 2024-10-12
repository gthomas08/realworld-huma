package user

import (
	"context"

	"github.com/gthomas08/realworld-huma/gen/postgres/public/model"
)

type Repository interface {
	GetUserByEmail(ctx context.Context, email string) (*model.Users, error)
	GetUserByEmailOrUsername(ctx context.Context, email string, username string) (*model.Users, error)
	CreateUser(ctx context.Context, user *model.Users) (*model.Users, error)
}
