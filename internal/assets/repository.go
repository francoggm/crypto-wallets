package assets

import (
	"context"
	"github/francoggm/crypto-wallets/internal/models"
)

type Repository interface {
	GetAllAssets(ctx context.Context) ([]*models.Asset, error)
	GetAssetByID(ctx context.Context, assetId int) (*models.Asset, error)
	GetAssetByName(ctx context.Context, assetName string) (*models.Asset, error)
	GetAssetTicker(ctx context.Context, assetId int) (*models.Ticker, error)
	GetAssetHistoricalData(ctx context.Context, assetId int, interval models.IntervalTime) ([]*models.Ticker, error)
	SaveAssetTicker(ctx context.Context, assetTicker *models.AssetTicker) error
}
