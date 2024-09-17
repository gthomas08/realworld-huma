package repository

import (
	"context"

	"github.com/gthomas08/realworld-huma/internal/db/postgres"
	"github.com/gthomas08/realworld-huma/internal/domain/user"

	. "github.com/go-jet/jet/v2/postgres"
	. "github.com/gthomas08/realworld-huma/internal/db/postgres/jet/postgres/public/table"

	"github.com/gthomas08/realworld-huma/internal/db/postgres/jet/postgres/public/model"
)

type repository struct {
	db *postgres.DB
}

func NewRepository(db *postgres.DB) user.Repository {
	return &repository{db: db}
}

func (r *repository) GetUser(ctx context.Context) (*model.Users, error) {
	var user model.Users

	stmt := SELECT(Users.AllColumns).
		FROM(Users).
		ORDER_BY(Users.ID.DESC())

	err := stmt.QueryContext(ctx, r.db.Conn, &user)

	return &user, err
}

func (r *repository) CreateUser(ctx context.Context, user *model.Users) error {
	insertStmt := Users.
		INSERT(Users.ID, Users.Username, Users.Email, Users.Password).
		MODEL(user)

	_, err := insertStmt.ExecContext(ctx, r.db.Conn)
	return err
}
