package user

import (
	"context"

	"github.com/gthomas08/realworld-huma/internal/domain/user/dtos"
)

type Usecase interface {
	CreateUser(ctx context.Context, input *dtos.CreateUserRequest) (*dtos.User, error)
	Login(ctx context.Context, input *dtos.LoginRequest) (*dtos.User, error)
}
