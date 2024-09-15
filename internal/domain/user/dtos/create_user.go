package dtos

type CreateUserRequest struct {
	Username string `json:"username" example:"johndoe" doc:"Username of the user"`
	Email    string `json:"email" example:"johndoe@example.com" doc:"Email of the user"`
	Password string `json:"password" example:"password" doc:"Password of the user"`
}

type CreateUserResponse struct {
	Id int64 `json:"id" example:"1" doc:"The id of the created user"`
}
