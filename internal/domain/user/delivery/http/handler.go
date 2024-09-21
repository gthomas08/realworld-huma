package http

import (
	"context"

	"github.com/gthomas08/realworld-huma/internal/domain/user"
	"github.com/gthomas08/realworld-huma/internal/domain/user/dtos"
	"github.com/gthomas08/realworld-huma/internal/utils/types"
	"github.com/gthomas08/realworld-huma/pkg/logger"
)

type userHandler struct {
	logger      *logger.Logger
	userUsecase user.Usecase
}

func NewHandler(logger *logger.Logger, userUsecase user.Usecase) *userHandler {
	return &userHandler{logger: logger, userUsecase: userUsecase}
}

type CreateUserResponse struct {
	User *dtos.User `json:"user" doc:"The created user"`
}

func (h *userHandler) CreateUser(ctx context.Context, input *types.RequestBody[dtos.CreateUserRequest]) (*types.ResponseBody[CreateUserResponse], error) {
	createdUser := h.userUsecase.CreateUser(ctx, &input.Body)

	resp := &types.ResponseBody[CreateUserResponse]{
		Body: CreateUserResponse{
			User: createdUser,
		},
	}

	return resp, nil
}
