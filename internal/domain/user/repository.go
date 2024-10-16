package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/gthomas08/realworld-huma/gen/postgres/public/model"
)

type Repository interface {
	GetUserById(ctx context.Context, id uuid.UUID) (*model.Users, error)
	GetUserByEmail(ctx context.Context, email string) (*model.Users, error)
	GetUserByEmailOrUsername(ctx context.Context, email string, username string) (*model.Users, error)

	CreateUser(ctx context.Context, user *model.Users) (*model.Users, error)
	UpdateUser(ctx context.Context, user *model.Users) (*model.Users, error)
}
