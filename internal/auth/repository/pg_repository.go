package repository

import (
	"context"
	"github/francoggm/crypto-wallets/internal/auth"
	"github/francoggm/crypto-wallets/internal/models"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
)

type authRepository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) auth.Repository {
	return &authRepository{
		db: db,
	}
}

func (repo *authRepository) FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "usecase.auth.FindUserByEmail")
	defer span.Finish()

	user := new(models.User)
	if err := repo.db.GetContext(ctx, user, getUserByEmail, email); err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *authRepository) FindUserByID(ctx context.Context, id int) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "usecase.auth.FindUserByID")
	defer span.Finish()

	user := new(models.User)
	if err := repo.db.GetContext(ctx, user, getUserByID, id); err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *authRepository) CreateUser(ctx context.Context, user *models.User) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "usecase.auth.CreateUser")
	defer span.Finish()

	rows, err := repo.db.NamedQueryContext(ctx, insertUser, &user)
	if err != nil {
		return err
	}

	if rows.Next() {
		rows.Scan(&user.ID)
	}

	return nil
}

func (repo *authRepository) UpdateLastLogin(ctx context.Context, userID int, lastLogin time.Time) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "usecase.auth.UpdateLastLogin")
	defer span.Finish()

	_, err := repo.db.ExecContext(ctx, updateLastLogin, userID, lastLogin)
	return err
}
