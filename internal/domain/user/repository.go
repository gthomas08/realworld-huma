package user

import (
	"context"

	"github.com/gthomas08/realworld-huma/internal/db/postgres/jet/postgres/public/model"
)

type Repository interface {
	CreateUser(ctx context.Context, user *model.Users) (*model.Users, error)
}
