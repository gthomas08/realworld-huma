package usecase

import (
	"context"

	"github.com/gthomas08/realworld-huma/internal/domain/ping"
)

type pingUsecase struct {
	pingRepository ping.Repository
}

func NewPingUsecase(pingRepository ping.Repository) ping.Usecase {
	return &pingUsecase{pingRepository: pingRepository}
}

func (uc *pingUsecase) GetPingMessage(ctx context.Context) string {
	res, err := uc.pingRepository.GetPingMessage(ctx)
	if err != nil {
		return "Entity not found"
	}
	return res
}
