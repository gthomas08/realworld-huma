package http

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
	"github.com/gthomas08/realworld-huma/internal/utils/security"
)

func (h *userHandler) RegisterRoutes(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "get-user",
		Method:      http.MethodGet,
		Path:        "/api/user",
		Summary:     "Gets current user",
		Security:    security.RequireAuth(security.Bearer),
	}, h.GetCurrentUser)

	huma.Register(api, huma.Operation{
		OperationID: "login-user",
		Method:      http.MethodPost,
		Path:        "/api/users/login",
		Summary:     "Logs in a user",
	}, h.Login)

	huma.Register(api, huma.Operation{
		OperationID: "register-user",
		Method:      http.MethodPost,
		Path:        "/api/users",
		Summary:     "Registers a new user",
	}, h.RegisterUser)

	huma.Register(api, huma.Operation{
		OperationID: "update-user",
		Method:      http.MethodPut,
		Path:        "/api/user",
		Summary:     "Updates current users email",
		Security:    security.RequireAuth(security.Bearer),
	}, h.UpdateUser)
}
