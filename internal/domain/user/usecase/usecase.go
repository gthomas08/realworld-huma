package usecase

import (
	"context"

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
	// Check if user already exists based on email or username
	if _, err := uc.userRepository.GetUserByEmailOrUsername(ctx, input.Email, input.Username); err == nil {
		return nil, errs.NewAppError(errs.ErrEntityExists)
	}

	user, err := uc.userRepository.CreateUser(ctx, mapper.CreateUserRequestToUserModel(input))
	if err != nil {
		uc.logger.Error("failed to create user", "error", err)
		return nil, errs.NewHTTPError(errs.ErrInternal)
	}

	return mapper.UserModelToUser(user), nil
}
