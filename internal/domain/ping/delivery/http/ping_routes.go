package http

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type PingResponse struct {
	Body struct {
		Message string `json:"message" example:"Pong!" doc:"Ping message"`
	}
}

// RegisterPingRoutes sets up routes related to ping operations.
func (ph *pingHandler) RegisterPingRoutes(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "ping",
		Method:      http.MethodGet,
		Path:        "/v1/ping",
		Summary:     "Gets ping message",
	},
		func(ctx context.Context, input *struct{}) (*PingResponse, error) {
			resp := PingResponse{}
			resp.Body.Message = ph.GetPingMessage(ctx, input)

			return &resp, nil
		})
}
