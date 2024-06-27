package wallets

import (
	"context"
	"github/francoggm/crypto-wallets/internal/models"
)

type Repository interface {
	GetUserWallets(ctx context.Context, userId int) ([]*models.Wallet, error)
	GetUserWallet(ctx context.Context, userId int, assetId int) (*models.Wallet, error)
	CreateWallet(ctx context.Context, userId int, assetId int) (*models.Wallet, error)
	UpdateWallet(ctx context.Context, quantity float64, wallet *models.Wallet) (*models.Wallet, error)
	UpdateWalletUSD(ctx context.Context, userId int, amount float64) error
	CreateTransaction(ctx context.Context, userId int, assetId int, transactionType string, amount float64, price float64) error
}
