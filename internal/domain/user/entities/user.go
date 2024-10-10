package entities

import (
	"github.com/gthomas08/realworld-huma/internal/db/postgres/jet/postgres/public/model"
	"github.com/gthomas08/realworld-huma/internal/domain/user/dtos"
)

func CreateUserRequestToUser(cur *dtos.RegisterUserRequest) *model.Users {
	return &model.Users{
		Username: cur.Username,
		Email:    cur.Email,
		Password: cur.Password,
	}
}
