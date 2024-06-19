package assets

import "github/francoggm/crypto-wallets/internal/models"

type Repository interface {
	GetAllAssets() ([]*models.Asset, error)
	GetAsset(assetName string) (*models.Asset, error)
	GetAssetTicker(asset models.Asset) (*models.AssetTicker, error)
}
