package dtos

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id" format:"uuid" doc:"The id of the user"`
	Username string    `json:"username" example:"johndoe" doc:"The username of the user"`
	Email    string    `json:"email" format:"email" doc:"The email of the user"`
	Bio      *string   `json:"bio" required:"false" doc:"The bio of the user"`
	Image    *string   `json:"image" required:"false" format:"uri" doc:"The image of the user"`
}

type RegisterUserRequest struct {
	Username string `json:"username" example:"johndoe" doc:"The username of the user"`
	Email    string `json:"email" format:"email" doc:"The email of the user"`
	Password string `json:"password" example:"password" doc:"Password of the user"`
}

type LoginRequest struct {
	Email    string `json:"email" format:"email" doc:"The email of the user"`
	Password string `json:"password" example:"password" doc:"Password of the user"`
}
