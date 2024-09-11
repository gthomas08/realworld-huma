package usecase

import (
	"context"

	"github.com/gthomas08/realworld-huma/internal/domain/ping"
)

type pingUsecase struct{}

func NewPingUsecase() ping.Usecase {
	return &pingUsecase{}
}

func (uc *pingUsecase) GetPingMessage(ctx context.Context) string {
	return "pong"
}
