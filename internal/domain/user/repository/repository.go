package repository

import (
	"github.com/gthomas08/realworld-huma/internal/db/sqlite"
	"github.com/gthomas08/realworld-huma/internal/domain/user"
)

type repository struct {
	db *sqlite.DB
}

func NewRepository(db *sqlite.DB) user.Repository {
	return &repository{db: db}
}
