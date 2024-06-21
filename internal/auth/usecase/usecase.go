package usecase

import (
	"context"
	"database/sql"
	"errors"
	"github/francoggm/crypto-wallets/internal/auth"
	"github/francoggm/crypto-wallets/internal/models"
	"github/francoggm/crypto-wallets/pkg/utils"
	"time"

	"github.com/gofiber/fiber/v3/log"
	"github.com/opentracing/opentracing-go"
)

type authUseCase struct {
	repo auth.Repository
}

func NewUseCase(repo auth.Repository) auth.UseCase {
	return &authUseCase{
		repo: repo,
	}
}

func (uc *authUseCase) Register(ctx context.Context, user *models.User) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "usecase.auth.Register")
	defer span.Finish()

	u, err := uc.repo.FindUserByEmail(ctx, user.Email)
	if u != nil {
		return auth.ErrEmailAlreadyRegistered
	} else if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Error(err)
		return auth.ErrFailedCreatingUser
	}

	if err := user.HashPassword(); err != nil {
		log.Error(err)
		return auth.ErrFailedCreatingUser
	}

	//default for now
	user.Role = "user"

	if err := uc.repo.CreateUser(ctx, user); err != nil {
		log.Error(err)
		return auth.ErrFailedCreatingUser
	}

	return nil
}

func (uc *authUseCase) Login(ctx context.Context, user *models.UserLogin) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "usecase.auth.Login")
	defer span.Finish()

	u, err := uc.repo.FindUserByEmail(ctx, user.Email)
	if u == nil || err != nil {
		log.Error(err)
		return nil, auth.ErrInvalidCredentials
	}

	if !utils.CompareHashPassword(u.Password, user.Password) {
		return nil, auth.ErrInvalidCredentials
	}

	return u, nil
}

func (uc *authUseCase) UpdateLastLogin(ctx context.Context, userID int, lastLogin time.Time) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "usecase.auth.UpdateLastLogin")
	defer span.Finish()

	return uc.repo.UpdateLastLogin(ctx, userID, lastLogin)
}
