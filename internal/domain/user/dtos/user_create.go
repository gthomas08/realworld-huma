package dtos

type CreateUserRequest struct {
	Body struct {
		Username string `json:"username" example:"johndoe" doc:"Username of the user"`
		Email    string `json:"email" example:"johndoe@example.com" doc:"Email of the user"`
	}
}
