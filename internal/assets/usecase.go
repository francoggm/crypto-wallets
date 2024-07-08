package assets

import (
	"context"
	"github/francoggm/crypto-wallets/internal/models"
)

type UseCase interface {
	GetAllAssetsData(ctx context.Context) ([]*models.AssetTicker, error)
	GetAssetData(ctx context.Context, assetName string) (*models.AssetTicker, error)
	GetAssetHistoricalData(ctx context.Context, assetName string, interval models.IntervalTime) (*models.AssetHistory, error)
}
