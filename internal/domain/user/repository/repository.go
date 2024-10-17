package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/gthomas08/realworld-huma/internal/db/postgres"
	"github.com/gthomas08/realworld-huma/internal/domain/user"
	"github.com/gthomas08/realworld-huma/pkg/errs"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/gthomas08/realworld-huma/gen/postgres/public/model"
	. "github.com/gthomas08/realworld-huma/gen/postgres/public/table"
)

type repository struct {
	db *postgres.DB
}

func NewRepository(db *postgres.DB) user.Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetUserById(ctx context.Context, id uuid.UUID) (*model.Users, error) {
	var user model.Users

	stmt := SELECT(Users.AllColumns).
		FROM(Users).
		WHERE(Users.ID.EQ(UUID(id)))

	err := stmt.QueryContext(ctx, r.db.Conn, &user)
	if err != nil {
		if errors.Is(err, qrm.ErrNoRows) {
			return nil, fmt.Errorf("user %w by id", errs.ErrNotFound)
		}
		return nil, err
	}

	return &user, err
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*model.Users, error) {
	var user model.Users

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

func (r *repository) GetUserByUsername(ctx context.Context, username string) (*model.Users, error) {
	var user model.Users

	stmt := SELECT(Users.AllColumns).
		FROM(Users).
		WHERE(Users.Username.EQ(String(username)))

	err := stmt.QueryContext(ctx, r.db.Conn, &user)
	if err != nil {
		if errors.Is(err, qrm.ErrNoRows) {
			return nil, fmt.Errorf("user %w by username", errs.ErrNotFound)
		}
		return nil, err
	}

	return &user, err
}

func (r *repository) GetUserByEmailOrUsername(ctx context.Context, email string, username string) (*model.Users, error) {
	var user model.Users

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

func (r *repository) CreateUser(ctx context.Context, user *model.Users) (*model.Users, error) {
	var newUser model.Users

	insertStmt := Users.
		INSERT(Users.AllColumns).
		MODEL(user).
		RETURNING(Users.AllColumns)

	err := insertStmt.QueryContext(ctx, r.db.Conn, &newUser)

	return &newUser, err
}

func (r *repository) UpdateUser(ctx context.Context, user *model.Users) (*model.Users, error) {
	var updatedUser model.Users

	updateStmt := Users.
		UPDATE(Users.AllColumns).
		MODEL(user).
		WHERE(Users.ID.EQ(UUID(user.ID))).
		RETURNING(Users.AllColumns)

	err := updateStmt.QueryContext(ctx, r.db.Conn, &updatedUser)

	return &updatedUser, err
}
