package entities

import (
	"github.com/gthomas08/realworld-huma/internal/domain/user/dtos"
	"github.com/gthomas08/realworld-huma/internal/domain/user/repository/sqlc"
)

func NewCreateUserParams(cur dtos.CreateUserRequest) sqlc.CreateUserParams {
	return sqlc.CreateUserParams{
		Email:    cur.Email,
		Username: cur.Username,
		Password: cur.Password,
	}
}
