package http

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
	"github.com/gthomas08/realworld-huma/internal/utils/openapi"
	"github.com/gthomas08/realworld-huma/internal/utils/security"
)

func (h *handler) RegisterRoutes(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "get-profile",
		Method:      http.MethodGet,
		Path:        "/profiles/{username}",
		Summary:     "Get a profile",
		Description: "Get a profile of a user of the system. Auth is optional.",
		Tags:        openapi.AddTags(openapi.ProfileTag),
		Security:    []map[string][]string{},
	}, h.GetProfile)

	huma.Register(api, huma.Operation{
		OperationID: "FollowUserByUsername",
		Method:      http.MethodGet,
		Path:        "/profiles/{username}/follow",
		Summary:     "Follow a user",
		Description: "Follow a user by username",
		Tags:        openapi.AddTags(openapi.ProfileTag),
		Security:    security.RequireAuth(security.Bearer),
	}, h.FollowUserByUsername)
}
