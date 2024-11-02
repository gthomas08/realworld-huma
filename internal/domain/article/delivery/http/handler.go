package http

import (
	"context"

	"github.com/gthomas08/realworld-huma/config"
	"github.com/gthomas08/realworld-huma/internal/domain/article"
	"github.com/gthomas08/realworld-huma/internal/domain/article/dtos"
	"github.com/gthomas08/realworld-huma/internal/utils/types"
	"github.com/gthomas08/realworld-huma/pkg/errs"
	"github.com/gthomas08/realworld-huma/pkg/logger"
)

type handler struct {
	cfg            *config.Config
	logger         *logger.Logger
	articleUsecase article.Usecase
}

func NewHandler(cfg *config.Config, logger *logger.Logger, articleUsecase article.Usecase) *handler {
	return &handler{
		cfg:            cfg,
		logger:         logger,
		articleUsecase: articleUsecase,
	}
}

type TagsResponse struct {
	Tags []dtos.Tag `json:"tags" nullable:"false" doc:"The list of available tags"`
}

func (h *handler) GetTags(ctx context.Context, input *struct{}) (*types.ResponseBody[TagsResponse], error) {
	tags, err := h.articleUsecase.GetTags(ctx)
	if err != nil {
		h.logger.Error("failed to get available tags", "error", err.Error())
		return nil, errs.ResolveError(err)
	}

	resp := &types.ResponseBody[TagsResponse]{
		Body: TagsResponse{
			Tags: tags,
		},
	}

	return resp, nil
}
