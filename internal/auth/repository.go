package auth

import (
	"context"
	"github/francoggm/crypto-wallets/internal/models"
	"time"
)

type Repository interface {
	FindUserByEmail(ctx context.Context, email string) (*models.User, error)
	FindUserByID(ctx context.Context, id int) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) error
	UpdateLastLogin(ctx context.Context, userID int, lastLogin time.Time) error
}
