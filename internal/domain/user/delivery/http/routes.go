package http

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
)

func (h *userHandler) RegisterRoutes(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "create-user",
		Method:      http.MethodPost,
		Path:        "/api/users",
		Summary:     "Registers a new user",
	}, h.CreateUser)

	huma.Register(api, huma.Operation{
		OperationID: "login-user",
		Method:      http.MethodPost,
		Path:        "/api/users/login",
		Summary:     "Logs in a user",
	}, h.Login)
}
