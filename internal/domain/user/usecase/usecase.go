package usecase

import (
	"context"
	"errors"

	"github.com/gthomas08/realworld-huma/internal/domain/user"
	"github.com/gthomas08/realworld-huma/internal/domain/user/dtos"
	"github.com/gthomas08/realworld-huma/internal/domain/user/mapper"
	"github.com/gthomas08/realworld-huma/pkg/errs"
	"github.com/gthomas08/realworld-huma/pkg/logger"
)

type userUsecase struct {
	logger         *logger.Logger
	userRepository user.Repository
}

func NewUsecase(logger *logger.Logger, userRepository user.Repository) user.Usecase {
	return &userUsecase{logger: logger, userRepository: userRepository}
}

func (uc *userUsecase) CreateUser(ctx context.Context, input *dtos.CreateUserRequest) (*dtos.User, error) {
	existingUser, err := uc.userRepository.GetUserByEmailOrUsername(ctx, input.Email, input.Username)
	if err != nil && !errors.Is(err, errs.ErrNotFound) {
		return nil, err
	}
	if existingUser != nil {
		return nil, errs.NewAppError(errs.EntityExists, "user with same email or username already exists")
	}

	newUser, err := uc.userRepository.CreateUser(ctx, mapper.CreateUserRequestToUserModel(input))
	if err != nil {
		return nil, err
	}

	return mapper.UserModelToUser(newUser), nil
}
