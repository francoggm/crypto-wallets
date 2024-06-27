package wallets

import (
	"context"
	"github/francoggm/crypto-wallets/internal/models"
)

type UseCase interface {
	GetUserWallets(ctx context.Context, userId int) ([]*models.WalletTicker, error)
	GetUserWallet(ctx context.Context, userId int, assetName string) (*models.WalletTicker, error)
	BuyAsset(ctx context.Context, user *models.User, transaction *models.TransactionAsset) (*models.WalletTicker, error)
	SellAsset(ctx context.Context, user *models.User, transaction *models.TransactionAsset) (*models.WalletTicker, error)
}
