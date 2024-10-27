package http

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
	"github.com/gthomas08/realworld-huma/internal/utils/openapi"
	"github.com/gthomas08/realworld-huma/internal/utils/security"
)

func (h *userHandler) RegisterRoutes(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID:   "Login",
		Method:        http.MethodPost,
		Path:          "/users/login",
		Summary:       "Existing user login",
		Description:   "Login for existing user",
		Tags:          openapi.AddTags(openapi.UserAndAuthTag),
		Security:      security.RequireNoAuth(),
		DefaultStatus: http.StatusOK,
	}, h.Login)

	huma.Register(api, huma.Operation{
		OperationID:   "CreateUser",
		Method:        http.MethodPost,
		Path:          "/users",
		Summary:       "Register a new user",
		Description:   "Register for new user",
		Tags:          openapi.AddTags(openapi.UserAndAuthTag),
		Security:      security.RequireNoAuth(),
		DefaultStatus: http.StatusCreated,
	}, h.RegisterUser)

	huma.Register(api, huma.Operation{
		OperationID:   "GetCurrentUser",
		Method:        http.MethodGet,
		Path:          "/user",
		Summary:       "Get current user",
		Description:   "Gets the currently logged-in user",
		Security:      security.RequireAuth(security.Bearer),
		Tags:          openapi.AddTags(openapi.UserAndAuthTag),
		DefaultStatus: http.StatusOK,
	}, h.GetCurrentUser)

	huma.Register(api, huma.Operation{
		OperationID:   "UpdateCurrentUser",
		Method:        http.MethodPut,
		Path:          "/user",
		Summary:       "Update current user",
		Description:   "Updated user information for current user",
		Security:      security.RequireAuth(security.Bearer),
		Tags:          openapi.AddTags(openapi.UserAndAuthTag),
		DefaultStatus: http.StatusOK,
	}, h.UpdateUser)
}
