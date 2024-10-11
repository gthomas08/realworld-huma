package user

import (
	"context"

	"github.com/gthomas08/realworld-huma/internal/domain/user/entities"
)

type Repository interface {
	GetUserByEmail(ctx context.Context, email string) (*entities.User, error)
	GetUserByEmailOrUsername(ctx context.Context, email string, username string) (*entities.User, error)
	CreateUser(ctx context.Context, user *entities.User) (*entities.User, error)
}
