package article

import (
	"context"

	"github.com/gthomas08/realworld-huma/gen/postgres/public/model"
)

type Repository interface {
	GetTags(ctx context.Context) ([]model.Tags, error)
}
