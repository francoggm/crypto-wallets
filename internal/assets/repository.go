package assets

import (
	"context"
	"github/francoggm/crypto-wallets/internal/models"
)

type Repository interface {
	GetAllAssets(ctx context.Context) ([]*models.Asset, error)
	GetAsset(ctx context.Context, assetName string) (*models.Asset, error)
	GetAssetTicker(ctx context.Context, asset *models.Asset) (*models.AssetTicker, error)
	SaveAssetTicker(ctx context.Context, assetTicker *models.AssetTicker) error
}
