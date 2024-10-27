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
		OperationID:   "GetProfileByUsername",
		Method:        http.MethodGet,
		Path:          "/profiles/{username}",
		Summary:       "Get a profile",
		Description:   "Get a profile of a user of the system. Auth is optional.",
		Tags:          openapi.AddTags(openapi.ProfileTag),
		Security:      security.RequireNoAuth(),
		DefaultStatus: http.StatusOK,
	}, h.GetProfile)

	huma.Register(api, huma.Operation{
		OperationID:   "FollowUserByUsername",
		Method:        http.MethodPost,
		Path:          "/profiles/{username}/follow",
		Summary:       "Follow a user",
		Description:   "Follow a user by username",
		Tags:          openapi.AddTags(openapi.ProfileTag),
		Security:      security.RequireAuth(security.Bearer),
		DefaultStatus: http.StatusCreated,
	}, h.FollowUserByUsername)

	huma.Register(api, huma.Operation{
		OperationID:   "UnfollowUserByUsername",
		Method:        http.MethodDelete,
		Path:          "/profiles/{username}/follow",
		Summary:       "Unfollow a user",
		Description:   "Unfollow a user by username",
		Tags:          openapi.AddTags(openapi.ProfileTag),
		Security:      security.RequireAuth(security.Bearer),
		DefaultStatus: http.StatusNoContent,
	}, h.UnfollowUserByUsername)
}
