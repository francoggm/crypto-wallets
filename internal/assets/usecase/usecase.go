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
		repo: repo,
	}
}

func (uc *assetsUseCase) GetAllAssetsData(ctx context.Context) ([]*models.AssetTicker, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "assets.usecase.GetAllAssetsData")
	defer span.Finish()

	assets, err := uc.repo.GetAllAssets(ctx)
	if err != nil {
		return nil, err
	}

	assetsTickers := make([]*models.AssetTicker, 0, len(assets))

	for _, asset := range assets {
		ticker, err := uc.repo.GetAssetTicker(ctx, asset.ID)
		if err != nil {
			log.Error(err)
			continue
		}

		var at models.AssetTicker
		at.Asset = *asset
		at.Ticker = *ticker

		assetsTickers = append(assetsTickers, &at)
	}

	return assetsTickers, nil
}

func (uc *assetsUseCase) GetAssetData(ctx context.Context, assetName string) (*models.AssetTicker, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "assets.usecase.GetAssetData")
	defer span.Finish()

	asset, err := uc.repo.GetAssetByName(ctx, assetName)
	if err != nil {
		return nil, err
	}

	ticker, err := uc.repo.GetAssetTicker(ctx, asset.ID)
	if err != nil {
		return nil, err
	}

	var at models.AssetTicker
	at.Asset = *asset
	at.Ticker = *ticker

	return &at, nil
}

func (uc *assetsUseCase) GetAssetHistoricalData(ctx context.Context, assetName string, interval models.IntervalTime) (*models.AssetHistory, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "assets.usecase.GetAssetHistoricalData")
	defer span.Finish()

	asset, err := uc.repo.GetAssetByName(ctx, assetName)
	if err != nil {
		return nil, err
	}

	tickers, err := uc.repo.GetAssetHistoricalData(ctx, asset.ID, interval)
	if err != nil {
		return nil, err
	}

	return &models.AssetHistory{
		Name:    asset.Name,
		Symbol:  asset.Symbol,
		Tickers: tickers,
	}, nil
}
