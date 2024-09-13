package http

import (
	"context"

	"github.com/gthomas08/realworld-huma/internal/domain/user"
	"github.com/gthomas08/realworld-huma/internal/domain/user/dtos"
	"github.com/gthomas08/realworld-huma/pkg/logger"
)

type userHandler struct {
	logger      *logger.Logger
	userUsecase user.Usecase
}

func NewHandler(logger *logger.Logger, userUsecase user.Usecase) *userHandler {
	return &userHandler{logger: logger, userUsecase: userUsecase}
}

func (uh *userHandler) CreateUser(ctx context.Context, input *dtos.CreateUserRequest) int {
	return 2
}
