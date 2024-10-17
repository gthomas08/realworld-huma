package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/gthomas08/realworld-huma/internal/db/postgres"
	"github.com/gthomas08/realworld-huma/internal/domain/profile"
	"github.com/gthomas08/realworld-huma/pkg/errs"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/gthomas08/realworld-huma/gen/postgres/public/model"
	. "github.com/gthomas08/realworld-huma/gen/postgres/public/table"
)

type repository struct {
	db *postgres.DB
}

func NewRepository(db *postgres.DB) profile.Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetFollow(ctx context.Context, followerID uuid.UUID, followeeID uuid.UUID) (*model.Follows, error) {

	var follows model.Follows

	stmt := SELECT(Follows.AllColumns).
		FROM(Follows).
		WHERE(AND(Follows.FollowerID.EQ(UUID(followerID)), Follows.FolloweeID.EQ(UUID(followeeID))))

	err := stmt.QueryContext(ctx, r.db.Conn, &follows)
	if err != nil {
		if errors.Is(err, qrm.ErrNoRows) {
			return nil, fmt.Errorf("follow %w", errs.ErrNotFound)
		}
		return nil, err
	}

	return &follows, err
}
