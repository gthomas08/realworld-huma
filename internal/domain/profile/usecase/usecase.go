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
