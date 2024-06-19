package assets

import "github/francoggm/crypto-wallets/internal/models"

type UseCase interface {
	ListAllAssetsData() ([]*models.AssetTicker, error)
	ListAssetData(assetName string) (*models.AssetTicker, error)
}
