package dtos

type Profile struct {
	Username  string  `json:"username" example:"johndoe" doc:"The username of the user"`
	Bio       *string `json:"bio" required:"false" doc:"The bio of the user"`
	Image     *string `json:"image" required:"false" format:"uri" doc:"The image of the user"`
	Following bool    `json:"following" doc:"The following status of the user"`
}
