package http

import (
	"context"

	"github.com/gthomas08/realworld-huma/internal/domain/ping"
	"github.com/gthomas08/realworld-huma/pkg/logger"
)

type pingHandler struct {
	logger      *logger.Logger
	pingUsecase ping.Usecase
}

func NewPingHandler(logger *logger.Logger, pingUsecase ping.Usecase) *pingHandler {
	return &pingHandler{logger: logger, pingUsecase: pingUsecase}
}

func (ph *pingHandler) GetPingMessage(ctx context.Context, input *struct{}) string {
	return ph.pingUsecase.GetPingMessage(ctx)
}
