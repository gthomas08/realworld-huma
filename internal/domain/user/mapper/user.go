package mapper

import (
	"github.com/google/uuid"
	"github.com/gthomas08/realworld-huma/internal/db/postgres/jet/postgres/public/model"
	"github.com/gthomas08/realworld-huma/internal/domain/user/dtos"
	"github.com/gthomas08/realworld-huma/internal/domain/user/entities"
)

func RegisterUserRequestToUser(user *dtos.RegisterUserRequest) *entities.User {
	return &entities.User{
		Users: &model.Users{
			ID:       uuid.New(),
			Username: user.Username,
			Email:    user.Email,
			Password: user.Password,
		},
	}
}

func UserToUser(user *entities.User) *dtos.User {
	return &dtos.User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Bio:      user.Bio,
		Image:    user.Image,
	}
}
