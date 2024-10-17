package http

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
	"github.com/gthomas08/realworld-huma/internal/utils/openapi"
)

func (h *handler) RegisterRoutes(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "get-profile",
		Method:      http.MethodGet,
		Path:        "/profile/{username}",
		Summary:     "Gets profile",
		Tags:        openapi.AddTags(openapi.ProfileTag),
	}, h.GetProfile)
}
