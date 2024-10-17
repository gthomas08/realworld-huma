package mapper

import (
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
