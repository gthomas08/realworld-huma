package mapper

import (
	"github.com/google/uuid"
	"github.com/gthomas08/realworld-huma/gen/postgres/public/model"
	"github.com/gthomas08/realworld-huma/internal/domain/user/dtos"
)

func RegisterUserRequestToUser(user *dtos.RegisterUserRequest) *model.Users {
	return &model.Users{
		ID:       uuid.New(),
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
}

func UserWithTokenToUser(user *model.Users, token string) *dtos.User {
	return &dtos.User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Bio:      user.Bio,
		Image:    user.Image,
		Token:    token,
	}
}
