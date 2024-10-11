package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/gthomas08/realworld-huma/internal/db/postgres"
	"github.com/gthomas08/realworld-huma/internal/domain/user"
	"github.com/gthomas08/realworld-huma/internal/domain/user/entities"
	"github.com/gthomas08/realworld-huma/pkg/errs"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	. "github.com/gthomas08/realworld-huma/internal/db/postgres/jet/postgres/public/table"
)

type repository struct {
	db *postgres.DB
}

func NewRepository(db *postgres.DB) user.Repository {
	return &repository{db: db}
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	var user entities.User

	stmt := SELECT(Users.AllColumns).
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

func (r *repository) GetUserByEmailOrUsername(ctx context.Context, email string, username string) (*entities.User, error) {
	var user entities.User

	stmt := SELECT(Users.AllColumns).
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

func (r *repository) CreateUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	var newUser entities.User

	insertStmt := Users.
		INSERT(Users.AllColumns).
		MODEL(user).
		RETURNING(Users.AllColumns)

	err := insertStmt.QueryContext(ctx, r.db.Conn, &newUser)

	return &newUser, err
}
