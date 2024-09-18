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

func (r *repository) CreateUser(ctx context.Context, user *model.Users) (*model.Users, error) {
	var dest model.Users

	insertStmt := Users.
		INSERT(Users.ID, Users.Username, Users.Email, Users.Password).
		MODEL(user).
		RETURNING(Users.ID, Users.Username, Users.Email, Users.Bio, Users.Image)

	err := insertStmt.QueryContext(ctx, r.db.Conn, &dest)

	return &dest, err
}
