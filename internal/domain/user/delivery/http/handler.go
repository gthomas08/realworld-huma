package http

import (
	"context"

	"github.com/gthomas08/realworld-huma/config"
	"github.com/gthomas08/realworld-huma/internal/domain/user"
	"github.com/gthomas08/realworld-huma/internal/domain/user/dtos"
	"github.com/gthomas08/realworld-huma/internal/utils/types"
	"github.com/gthomas08/realworld-huma/pkg/errs"
	"github.com/gthomas08/realworld-huma/pkg/logger"
)

type userHandler struct {
	cfg         *config.Config
	logger      *logger.Logger
	userUsecase user.Usecase
}

func NewHandler(cfg *config.Config, logger *logger.Logger, userUsecase user.Usecase) *userHandler {
	return &userHandler{
		cfg:         cfg,
		logger:      logger,
		userUsecase: userUsecase,
	}
}

type UserResponse struct {
	User *dtos.User `json:"user" doc:"The user"`
}

func (h *userHandler) Login(ctx context.Context, input *types.RequestBody[dtos.LoginRequest]) (*types.ResponseBody[UserResponse], error) {
	user, err := h.userUsecase.Login(ctx, &input.Body)
	if err != nil {
		h.logger.Error("failed to login user", "error", err.Error())
		return nil, errs.ResolveError(err)
	}

	resp := &types.ResponseBody[UserResponse]{
		Body: UserResponse{
			User: user,
		},
	}

	return resp, nil
}

func (h *userHandler) RegisterUser(ctx context.Context, input *types.RequestBody[dtos.RegisterUserRequest]) (*types.ResponseBody[UserResponse], error) {
	user, err := h.userUsecase.RegisterUser(ctx, &input.Body)
	if err != nil {
		h.logger.Error("failed to create user", "error", err.Error())
		return nil, errs.ResolveError(err)
	}

	resp := &types.ResponseBody[UserResponse]{
		Body: UserResponse{
			User: user,
		},
	}

	return resp, nil
}

func (h *userHandler) GetCurrentUser(ctx context.Context, input *struct{}) (*types.ResponseBody[UserResponse], error) {
	user, err := h.userUsecase.GetCurrentUser(ctx)
	if err != nil {
		h.logger.Error("failed to get current user", "error", err.Error())
		return nil, errs.ResolveError(err)
	}

	resp := &types.ResponseBody[UserResponse]{
		Body: UserResponse{
			User: user,
		},
	}

	return resp, nil
}

func (h *userHandler) UpdateUser(ctx context.Context, input *types.RequestBody[dtos.UpdateUserRequest]) (*types.ResponseBody[UserResponse], error) {
	user, err := h.userUsecase.UpdateUser(ctx, &input.Body)
	if err != nil {
		h.logger.Error("failed to update users email", "error", err.Error())
		return nil, errs.ResolveError(err)
	}

	resp := &types.ResponseBody[UserResponse]{
		Body: UserResponse{
			User: user,
		},
	}

	return resp, nil
}
