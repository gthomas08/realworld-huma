package mapper

import (
	"github.com/google/uuid"
	"github.com/gthomas08/realworld-huma/gen/postgres/public/model"
	"github.com/gthomas08/realworld-huma/internal/domain/profile/dtos"
)

func UserWithFollowingToProfile(user *model.Users, following bool) *dtos.Profile {
	return &dtos.Profile{
		Username:  user.Username,
		Bio:       user.Bio,
		Image:     user.Image,
		Following: following,
	}
}

func NewFollow(followerID uuid.UUID, followeeID uuid.UUID) *model.Follows {
	return &model.Follows{
		ID:         uuid.New(),
		FollowerID: followerID,
		FolloweeID: followeeID,
	}
}
