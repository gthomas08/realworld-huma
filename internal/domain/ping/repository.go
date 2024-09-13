package ping

import (
	"context"
	"database/sql"

	"github.com/gthomas08/realworld-huma/internal/domain/ping/repository/sqlc"
)

type Repository interface {
	Atomic(ctx context.Context, opt *sql.TxOptions, repo func(tx Repository) error) error
	sqlc.Querier
}
