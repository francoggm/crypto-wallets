package auth

import (
	"context"
	"github/francoggm/crypto-wallets/internal/models"
	"time"
)

type UseCase interface {
	Register(ctx context.Context, user *models.User) error
	Login(ctx context.Context, user *models.UserLogin) (*models.User, error)
	UpdateLastLogin(ctx context.Context, userID int, lastLogin time.Time) error
}
