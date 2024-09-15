package usecase

import (
	"context"

	"github.com/gthomas08/realworld-huma/internal/domain/user"
	"github.com/gthomas08/realworld-huma/internal/domain/user/dtos"
	"github.com/gthomas08/realworld-huma/internal/domain/user/entities"
)

type userUsecase struct {
	userRepository user.Repository
}

func NewUsecase(userRepository user.Repository) user.Usecase {
	return &userUsecase{userRepository: userRepository}
}

func (uc *userUsecase) CreateUser(ctx context.Context, input *dtos.CreateUserRequest) int64 {
	user := entities.NewCreateUserParams(*input)

	id, err := uc.userRepository.CreateUser(ctx, user)

	if err != nil {
		return -1
	}

	return id
}
