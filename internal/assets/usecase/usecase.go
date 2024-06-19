package usecase

import (
	"github/francoggm/crypto-wallets/internal/assets"
	"github/francoggm/crypto-wallets/internal/models"

	"github.com/gofiber/fiber/v3/log"
)

type assetsUseCase struct {
	repo assets.Repository
}

func NewUseCase(repo assets.Repository) assets.UseCase {
	return &assetsUseCase{
		repo,
	}
}

func (uc *assetsUseCase) ListAllAssetsData() ([]*models.AssetTicker, error) {
	assets, err := uc.repo.GetAllAssets()
	if err != nil {
		return nil, err
	}

	assetsTickers := make([]*models.AssetTicker, 0, len(assets))

	for _, asset := range assets {
		assetTicker, err := uc.repo.GetAssetTicker(*asset)
		if err != nil {
			log.Error(err)
			continue
		}

		assetsTickers = append(assetsTickers, assetTicker)
	}

	return assetsTickers, nil
}

func (uc *assetsUseCase) ListAssetData(assetName string) (*models.AssetTicker, error) {
	asset, err := uc.repo.GetAsset(assetName)
	if err != nil {
		return nil, err
	}

	assetTicker, err := uc.repo.GetAssetTicker(*asset)
	if err != nil {
		return nil, err
	}

	return assetTicker, nil
}
