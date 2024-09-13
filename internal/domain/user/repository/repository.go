package repository

import (
	"context"
	"database/sql"

	"github.com/gthomas08/realworld-huma/internal/db/sqlite"
	"github.com/gthomas08/realworld-huma/internal/domain/user"
	"github.com/gthomas08/realworld-huma/internal/domain/user/repository/sqlc"
)

type repository struct {
	db *sqlite.DB
	*sqlc.Queries
}

func NewRepository(db *sqlite.DB) user.Repository {
	return &repository{db: db, Queries: sqlc.New(db.Conn)}
}

// Atomic implements Repository Interface for transaction query
func (r *repository) Atomic(ctx context.Context, opt *sql.TxOptions, repo func(tx user.Repository) error) error {
	tx, err := r.db.Conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	qtx := r.Queries.WithTx(tx)

	newUserRepository := &repository{db: r.db, Queries: qtx}

	repo(newUserRepository)

	return tx.Commit()
}
