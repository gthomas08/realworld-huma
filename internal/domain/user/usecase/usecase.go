package usecase

import (
	"context"

	"github.com/gthomas08/realworld-huma/internal/domain/user"
	"github.com/gthomas08/realworld-huma/internal/domain/user/dtos"
	"github.com/gthomas08/realworld-huma/internal/domain/user/mapper"
)

type userUsecase struct {
	userRepository user.Repository
}

func NewUsecase(userRepository user.Repository) user.Usecase {
	return &userUsecase{userRepository: userRepository}
}

func (uc *userUsecase) CreateUser(ctx context.Context, input *dtos.CreateUserRequest) *dtos.User {
	user, _ := uc.userRepository.CreateUser(ctx, mapper.CreateUserRequestToUserModel(input))

	return mapper.UserModelToUser(user)
}
