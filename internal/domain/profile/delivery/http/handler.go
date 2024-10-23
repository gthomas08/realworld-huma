package http

import (
	"context"

	"github.com/gthomas08/realworld-huma/config"
	"github.com/gthomas08/realworld-huma/internal/domain/profile"
	"github.com/gthomas08/realworld-huma/internal/domain/profile/dtos"
	"github.com/gthomas08/realworld-huma/internal/utils/types"
	"github.com/gthomas08/realworld-huma/pkg/errs"
	"github.com/gthomas08/realworld-huma/pkg/logger"
)

type handler struct {
	cfg            *config.Config
	logger         *logger.Logger
	profileUsecase profile.Usecase
}

func NewHandler(cfg *config.Config, logger *logger.Logger, profileUsecase profile.Usecase) *handler {
	return &handler{
		cfg:            cfg,
		logger:         logger,
		profileUsecase: profileUsecase,
	}
}

type ProfileResponse struct {
	Profile *dtos.Profile `json:"profile" doc:"The profile of the user"`
}

func (h *handler) GetProfile(ctx context.Context, input *struct {
	Username string `path:"username"`
}) (*types.ResponseBody[ProfileResponse], error) {
	profile, err := h.profileUsecase.GetProfile(ctx, input.Username)
	if err != nil {
		h.logger.Error("failed to get profile", "error", err.Error())
		return nil, errs.ResolveError(err)
	}

	resp := &types.ResponseBody[ProfileResponse]{
		Body: ProfileResponse{
			Profile: profile,
		},
	}

	return resp, nil
}

func (h *handler) FollowUserByUsername(ctx context.Context, input *struct {
	Username string `path:"username"`
}) (*types.ResponseBody[ProfileResponse], error) {
	profile, err := h.profileUsecase.FollowUserByUsername(ctx, input.Username)
	if err != nil {
		h.logger.Error("failed to follow user", "error", err.Error())
		return nil, errs.ResolveError(err)
	}

	resp := &types.ResponseBody[ProfileResponse]{
		Body: ProfileResponse{
			Profile: profile,
		},
	}

	return resp, nil
}
