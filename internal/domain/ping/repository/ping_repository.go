package repository

import (
	"context"
	"database/sql"

	"github.com/gthomas08/realworld-huma/internal/db/sqlite"
	"github.com/gthomas08/realworld-huma/internal/domain/ping"
	"github.com/gthomas08/realworld-huma/internal/domain/ping/repository/sqlc"
)

type repository struct {
	db *sqlite.DB
	*sqlc.Queries
}

func NewPingRepository(db *sqlite.DB) ping.Repository {
	return &repository{db: db, Queries: sqlc.New(db.Conn)}
}

// Atomic implements Repository Interface for transaction query
func (r *repository) Atomic(ctx context.Context, opt *sql.TxOptions, repo func(tx ping.Repository) error) error {
	tx, err := r.db.Conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	qtx := r.Queries.WithTx(tx)

	newPingRepository := &repository{db: r.db, Queries: qtx}

	repo(newPingRepository)

	return tx.Commit()
}
