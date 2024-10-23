package usecase

import (
	"context"
	"errors"

	"github.com/gthomas08/realworld-huma/config"
	"github.com/gthomas08/realworld-huma/internal/ctxkit"
	"github.com/gthomas08/realworld-huma/internal/domain/profile"
	"github.com/gthomas08/realworld-huma/internal/domain/profile/dtos"
	"github.com/gthomas08/realworld-huma/internal/domain/profile/mapper"
	"github.com/gthomas08/realworld-huma/internal/domain/user"
	"github.com/gthomas08/realworld-huma/pkg/errs"
	"github.com/gthomas08/realworld-huma/pkg/logger"
)

type Usecase struct {
	cfg               *config.Config
	logger            *logger.Logger
	profileRepository profile.Repository
	userRepository    user.Repository
}

func NewUsecase(cfg *config.Config, logger *logger.Logger, profileRepository profile.Repository, userRepository user.Repository) profile.Usecase {
	return &Usecase{
		cfg:               cfg,
		logger:            logger,
		profileRepository: profileRepository,
		userRepository:    userRepository,
	}
}

func (u *Usecase) FollowUserByUsername(ctx context.Context, username string) (*dtos.Profile, error) {
	// Get the user claim from context
	userClaim, exists := ctxkit.GetUserContext(ctx)
	if !exists {
		return nil, errs.NewAppError(errs.EntityNotFound, "user claim not found")
	}

	// Get the user profile
	user, err := u.userRepository.GetUserByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return nil, errs.NewAppError(errs.EntityNotFound, "user not found")
		}
		return nil, err
	}

	// If the logged in user is the same as the one to follow, return an error
	if userClaim.ID == user.ID {
		return nil, errs.NewAppError(errs.InvalidOperation, "user cannot follow theirself")
	}

	// Create a new follow
	newFollow := mapper.NewFollow(userClaim.ID, user.ID)

	// Follow the user
	_, err = u.profileRepository.CreateFollow(ctx, newFollow)
	if err != nil {
		return nil, err
	}

	// Get the updated profile
	profile, err := u.GetProfile(ctx, username)
	if err != nil {
		return nil, err
	}

	return profile, nil
}

func (u *Usecase) GetProfile(ctx context.Context, username string) (*dtos.Profile, error) {
	// Get the user profile
	user, err := u.userRepository.GetUserByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return nil, errs.NewAppError(errs.EntityNotFound, "user not found")
		}
		return nil, err
	}

	// Determine if the user is following the logged in user
	isFollowing := false

	// Get the user claim from context
	userClaim, exists := ctxkit.GetUserContext(ctx)
	if exists {
		follow, err := u.profileRepository.GetFollow(ctx, userClaim.ID, user.ID)
		if err != nil {
			return nil, err
		}
		if follow != nil {
			isFollowing = true
		}
	}

	// Map the user profile with the following status
	return mapper.UserWithFollowingToProfile(user, isFollowing), nil
}
