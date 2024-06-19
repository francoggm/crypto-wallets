package utils

import (
	"github/francoggm/crypto-wallets/internal/models"
	"strings"
)

func GetAssetsListQueryParam(assets []*models.Asset) string {
	var assetsName []string
	for _, asset := range assets {
		assetsName = append(assetsName, asset.SearchName)
	}

	return strings.Join(assetsName, ",")
}

func GetAssetFromAssets(assets []*models.Asset, assetSearchName, assetName, assetSymbol string) *models.Asset {
	for _, asset := range assets {
		if asset.SearchName == assetSearchName || asset.Name == assetName || asset.Symbol == assetSymbol {
			return asset
		}
	}

	return nil
}
