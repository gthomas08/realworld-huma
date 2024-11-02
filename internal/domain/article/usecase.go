package article

import (
	"context"

	"github.com/gthomas08/realworld-huma/internal/domain/article/dtos"
)

type Usecase interface {
	GetTags(ctx context.Context) ([]dtos.Tag, error)
}
