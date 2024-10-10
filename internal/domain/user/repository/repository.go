package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/gthomas08/realworld-huma/internal/db/postgres"
	"github.com/gthomas08/realworld-huma/internal/domain/user"
	"github.com/gthomas08/realworld-huma/pkg/errs"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	. "github.com/gthomas08/realworld-huma/internal/db/postgres/jet/postgres/public/table"

	"github.com/gthomas08/realworld-huma/internal/db/postgres/jet/postgres/public/model"
)

type repository struct {
	db *postgres.DB
}

func NewRepository(db *postgres.DB) user.Repository {
	return &repository{db: db}
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*model.Users, error) {
	var user model.Users

	stmt := SELECT(Users.ID, Users.Username, Users.Email, Users.Password, Users.Bio, Users.Image).
		FROM(Users).
		WHERE(Users.Email.EQ(String(email)))

	err := stmt.QueryContext(ctx, r.db.Conn, &user)
	if err != nil {
		if errors.Is(err, qrm.ErrNoRows) {
			return nil, fmt.Errorf("user %w by email", errs.ErrNotFound)
		}
		return nil, err
	}

	return &user, err
}

func (r *repository) GetUserByEmailOrUsername(ctx context.Context, email string, username string) (*model.Users, error) {
	var user model.Users

	stmt := SELECT(Users.ID, Users.Username, Users.Email, Users.Bio, Users.Image).
		FROM(Users).
		WHERE(OR(Users.Email.EQ(String(email)), Users.Username.EQ(String(username))))

	err := stmt.QueryContext(ctx, r.db.Conn, &user)
	if err != nil {
		if errors.Is(err, qrm.ErrNoRows) {
			return nil, fmt.Errorf("user %w by email or username", errs.ErrNotFound)
		}
		return nil, err
	}

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
