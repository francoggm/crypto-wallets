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

func GetIntervalTime(interval string) models.IntervalTime {
	switch interval {
	case "7d":
		return models.IntervalTime7d
	case "30d":
		return models.IntervalTime30d
	case "90d":
		return models.IntervalTime90d
	case "1y":
		return models.IntervalTime1y
	case "max":
		return models.IntervalTimeMax
	default:
		return models.IntervalTime24h
	}
}

func GetIntervalTimeString(interval models.IntervalTime) string {
	switch interval {
	case models.IntervalTime7d:
		return "7d"
	case models.IntervalTime30d:
		return "30d"
	case models.IntervalTime90d:
		return "90d"
	case models.IntervalTime1y:
		return "1y"
	case models.IntervalTimeMax:
		return "max"
	default:
		return "24h"
	}
}
