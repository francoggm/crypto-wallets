package usecase

import (
	"context"
	"github/francoggm/crypto-wallets/internal/assets"
	"github/francoggm/crypto-wallets/internal/models"

	"github.com/gofiber/fiber/v3/log"
	"github.com/opentracing/opentracing-go"
)

type assetsUseCase struct {
	repo assets.Repository
}

func NewUseCase(repo assets.Repository) assets.UseCase {
	return &assetsUseCase{
		repo,
	}
}

func (uc *assetsUseCase) ListAllAssetsData(ctx context.Context) ([]*models.AssetTicker, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "assets.usecase.ListAllAssetsData")
	defer span.Finish()

	assets, err := uc.repo.GetAllAssets(ctx)
	if err != nil {
		return nil, err
	}

	assetsTickers := make([]*models.AssetTicker, 0, len(assets))

	for _, asset := range assets {
		assetTicker, err := uc.repo.GetAssetTicker(ctx, asset)
		if err != nil {
			log.Error(err)
			continue
		}

		assetsTickers = append(assetsTickers, assetTicker)
	}

	return assetsTickers, nil
}

func (uc *assetsUseCase) ListAssetData(ctx context.Context, assetName string) (*models.AssetTicker, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "assets.usecase.ListAssetData")
	defer span.Finish()

	asset, err := uc.repo.GetAsset(ctx, assetName)
	if err != nil {
		return nil, err
	}

	assetTicker, err := uc.repo.GetAssetTicker(ctx, asset)
	if err != nil {
		return nil, err
	}

	return assetTicker, nil
}
