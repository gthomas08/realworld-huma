package usecase

import (
	"context"

	"github.com/gthomas08/realworld-huma/config"
	"github.com/gthomas08/realworld-huma/internal/domain/article"
	"github.com/gthomas08/realworld-huma/internal/domain/article/dtos"
	"github.com/gthomas08/realworld-huma/internal/domain/article/mapper"
	"github.com/gthomas08/realworld-huma/pkg/logger"
)

type usecase struct {
	cfg               *config.Config
	logger            *logger.Logger
	articleRepository article.Repository
}

func NewUsecase(cfg *config.Config, logger *logger.Logger, articleRepository article.Repository) article.Usecase {
	return &usecase{
		cfg:               cfg,
		logger:            logger,
		articleRepository: articleRepository,
	}
}

func (uc *usecase) GetTags(ctx context.Context) ([]dtos.Tag, error) {
	tags, err := uc.articleRepository.GetTags(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.TagsToTags(tags), nil
}
