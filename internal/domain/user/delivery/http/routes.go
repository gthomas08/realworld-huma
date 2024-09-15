package http

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
	"github.com/gthomas08/realworld-huma/internal/domain/user/dtos"
)

type CreateUserResponse struct {
	Body struct {
		Id int64 `json:"message" example:"2" doc:"Ping message"`
	}
}

func (uh *userHandler) RegisterRoutes(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "create-user",
		Method:      http.MethodPost,
		Path:        "/api/users",
		Summary:     "Register a new user",
	}, func(ctx context.Context, input *dtos.CreateUserRequest) (*CreateUserResponse, error) {
		resp := CreateUserResponse{}
		resp.Body.Id = uh.CreateUser(ctx, input)

		return &resp, nil
	})
}
