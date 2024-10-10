package user

import (
	"context"

	"github.com/gthomas08/realworld-huma/internal/domain/user/dtos"
)

type Usecase interface {
	Login(ctx context.Context, input *dtos.LoginRequest) (*dtos.User, error)
	RegisterUser(ctx context.Context, input *dtos.RegisterUserRequest) (*dtos.User, error)
}
