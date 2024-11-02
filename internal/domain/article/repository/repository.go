package repository

import (
	"context"

	"github.com/gthomas08/realworld-huma/internal/db/postgres"
	"github.com/gthomas08/realworld-huma/internal/domain/article"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/gthomas08/realworld-huma/gen/postgres/public/model"
	. "github.com/gthomas08/realworld-huma/gen/postgres/public/table"
)

type repository struct {
	db *postgres.DB
}

func NewRepository(db *postgres.DB) article.Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetTags(ctx context.Context) ([]model.Tags, error) {
	var tags []model.Tags

	stmt := SELECT(Tags.AllColumns).
		FROM(Tags)

	err := stmt.QueryContext(ctx, r.db.Conn, &tags)

	return tags, err
}
