package usecase

import (
	"context"
	"errors"

	"github.com/gthomas08/realworld-huma/config"
	"github.com/gthomas08/realworld-huma/internal/domain/user"
	"github.com/gthomas08/realworld-huma/internal/domain/user/dtos"
	"github.com/gthomas08/realworld-huma/internal/domain/user/mapper"
	"github.com/gthomas08/realworld-huma/pkg/crypt"
	"github.com/gthomas08/realworld-huma/pkg/errs"
	"github.com/gthomas08/realworld-huma/pkg/jwtkit"
	"github.com/gthomas08/realworld-huma/pkg/logger"
)

type userUsecase struct {
	cfg            *config.Config
	logger         *logger.Logger
	userRepository user.Repository
}

func NewUsecase(cfg *config.Config, logger *logger.Logger, userRepository user.Repository) user.Usecase {
	return &userUsecase{
		cfg:            cfg,
		logger:         logger,
		userRepository: userRepository,
	}
}

func (uc *userUsecase) Login(ctx context.Context, input *dtos.LoginRequest) (*dtos.User, error) {
	user, err := uc.userRepository.GetUserByEmail(ctx, input.Email)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return nil, errs.NewAppError(errs.InvalidCredentials, "invalid email")
		}
		return nil, err
	}

	if !crypt.CheckPasswordHash(input.Password, user.Password) {
		return nil, errs.NewAppError(errs.InvalidCredentials, "invalid password")
	}

	token, err := jwtkit.GenerateToken(uc.cfg.App.Name, user.ID.String(), uc.cfg.JWT.Key, uc.cfg.JWT.Expired)
	if err != nil {
		return nil, err
	}

	return mapper.UserWithTokenToUser(user, token), nil
}

func (uc *userUsecase) RegisterUser(ctx context.Context, input *dtos.RegisterUserRequest) (*dtos.User, error) {
	existingUser, err := uc.userRepository.GetUserByEmailOrUsername(ctx, input.Email, input.Username)
	if err != nil && !errors.Is(err, errs.ErrNotFound) {
		return nil, err
	}
	if existingUser != nil {
		return nil, errs.NewAppError(errs.EntityExists, "user with same email or username already exists")
	}

	input.Password, err = crypt.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	newUser, err := uc.userRepository.CreateUser(ctx, mapper.RegisterUserRequestToUser(input))
	if err != nil {
		return nil, err
	}

	token, err := jwtkit.GenerateToken(uc.cfg.App.Name, newUser.ID.String(), uc.cfg.JWT.Key, uc.cfg.JWT.Expired)
	if err != nil {
		return nil, err
	}

	return mapper.UserWithTokenToUser(newUser, token), nil
}
