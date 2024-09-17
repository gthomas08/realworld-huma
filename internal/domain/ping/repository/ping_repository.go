package repository

import (
	"github.com/gthomas08/realworld-huma/internal/db/sqlite"
	"github.com/gthomas08/realworld-huma/internal/domain/ping"
)

type repository struct {
	db *sqlite.DB
}

func NewPingRepository(db *sqlite.DB) ping.Repository {
	return &repository{db: db}
}
