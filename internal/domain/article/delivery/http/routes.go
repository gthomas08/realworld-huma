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
		OperationID:   "GetTags",
		Method:        http.MethodGet,
		Path:          "/tags",
		Summary:       "Get tags",
		Description:   "Get tags. Auth not required.",
		Security:      security.Public(),
		Tags:          openapi.AddTags(openapi.TagsTag),
		DefaultStatus: http.StatusOK,
	}, h.GetTags)
}
