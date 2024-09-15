package repository

import (
	"context"
	"database/sql"
	"log"

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

// Atomic implements Repository Interface for transaction query.
func (r *repository) Atomic(ctx context.Context, opt *sql.TxOptions, repo func(tx user.Repository) error) error {
	tx, err := r.db.Conn.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err := tx.Rollback(); err != nil && err != sql.ErrTxDone {
			log.Printf("error rolling back transaction: %v", err)
		}
	}()
	qtx := r.Queries.WithTx(tx)

	newUserRepository := &repository{db: r.db, Queries: qtx}
	if err := repo(newUserRepository); err != nil {
		return err
	}

	return tx.Commit()
}
