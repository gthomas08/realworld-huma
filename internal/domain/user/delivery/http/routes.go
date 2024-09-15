package http

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
	"github.com/gthomas08/realworld-huma/internal/domain/user/dtos"
	"github.com/gthomas08/realworld-huma/internal/utils/types"
)

func (h *userHandler) RegisterRoutes(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "create-user",
		Method:      http.MethodPost,
		Path:        "/api/users",
		Summary:     "Registers a new user",
	}, func(ctx context.Context, input *types.RequestBody[dtos.CreateUserRequest]) (*types.ResponseBody[dtos.CreateUserResponse], error) {
		var resp types.ResponseBody[dtos.CreateUserResponse]

		resp.Body.Id = h.CreateUser(ctx, &input.Body)

		return &resp, nil
	})
}
