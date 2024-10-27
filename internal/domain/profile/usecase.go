package profile

import (
	"context"

	"github.com/gthomas08/realworld-huma/internal/domain/profile/dtos"
)

type Usecase interface {
	FollowUserByUsername(ctx context.Context, username string) (*dtos.Profile, error)
	UnfollowUserByUsername(ctx context.Context, username string) (*dtos.Profile, error)

	GetProfile(ctx context.Context, username string) (*dtos.Profile, error)
}
