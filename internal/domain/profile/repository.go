package profile

import (
	"context"

	"github.com/google/uuid"
	"github.com/gthomas08/realworld-huma/gen/postgres/public/model"
)

type Repository interface {
	CreateFollow(ctx context.Context, follows *model.Follows) (*model.Follows, error)

	GetFollow(ctx context.Context, followerID uuid.UUID, followeeID uuid.UUID) (*model.Follows, error)
}
