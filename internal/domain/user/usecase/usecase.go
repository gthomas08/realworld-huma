package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/gthomas08/realworld-huma/internal/db/postgres/jet/postgres/public/model"
	"github.com/gthomas08/realworld-huma/internal/domain/user"
	"github.com/gthomas08/realworld-huma/internal/domain/user/dtos"
)

type userUsecase struct {
	userRepository user.Repository
}

func NewUsecase(userRepository user.Repository) user.Usecase {
	return &userUsecase{userRepository: userRepository}
}

func (uc *userUsecase) CreateUser(ctx context.Context, input *dtos.CreateUserRequest) int64 {
	uc.userRepository.CreateUser(ctx, &model.Users{
		ID:       uuid.New(),
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	})
	return -2
}
