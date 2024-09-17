package repository

import (
	"github.com/gthomas08/realworld-huma/internal/db/postgres"
	"github.com/gthomas08/realworld-huma/internal/domain/ping"
)

type repository struct {
	db *postgres.DB
}

func NewPingRepository(db *postgres.DB) ping.Repository {
	return &repository{db: db}
}
