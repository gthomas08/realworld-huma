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

type UserResponse struct {
	User *dtos.User `json:"user" doc:"The created user"`
}

func (h *userHandler) CreateUser(ctx context.Context, input *types.RequestBody[dtos.CreateUserRequest]) (*types.ResponseBody[UserResponse], error) {
	var resp types.ResponseBody[UserResponse]

	user := UserResponse{
		User: h.userUsecase.CreateUser(ctx, &input.Body),
	}

	resp.Body = user

	return &resp, nil
}
