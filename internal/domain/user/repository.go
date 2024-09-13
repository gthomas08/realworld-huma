package user

import (
	"context"
	"database/sql"

	"github.com/gthomas08/realworld-huma/internal/domain/user/repository/sqlc"
)

type Repository interface {
	sqlc.Querier
	Atomic(ctx context.Context, opt *sql.TxOptions, repo func(tx Repository) error) error
}
