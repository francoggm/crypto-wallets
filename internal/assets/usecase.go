package assets

import (
	"context"
	"github/francoggm/crypto-wallets/internal/models"
)

type UseCase interface {
	ListAllAssetsData(ctx context.Context) ([]*models.AssetTicker, error)
	ListAssetData(ctx context.Context, assetName string) (*models.AssetTicker, error)
}
