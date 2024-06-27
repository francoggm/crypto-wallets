package repository

import (
	"context"
	"github/francoggm/crypto-wallets/internal/assets"
	"github/francoggm/crypto-wallets/internal/models"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
)

type assetsRepository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) assets.Repository {
	return &assetsRepository{
		db: db,
	}
}

func (r *assetsRepository) GetAllAssets(ctx context.Context) ([]*models.Asset, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "assets.repository.GetAllAssets")
	defer span.Finish()

	var assets []*models.Asset

	if err := r.db.SelectContext(ctx, &assets, getAllAssets); err != nil {
		return nil, err
	}

	return assets, nil
}

func (r *assetsRepository) GetAssetByID(ctx context.Context, assetId int) (*models.Asset, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "assets.repository.GetAsset")
	defer span.Finish()

	var asset models.Asset

	if err := r.db.GetContext(ctx, &asset, getAssetByID, assetId); err != nil {
		return nil, err
	}

	return &asset, nil
}

func (r *assetsRepository) GetAssetByName(ctx context.Context, assetName string) (*models.Asset, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "assets.repository.GetAsset")
	defer span.Finish()

	var asset models.Asset

	if err := r.db.GetContext(ctx, &asset, getAssetByName, assetName, strings.ToLower(assetName), strings.ToUpper(assetName)); err != nil {
		return nil, err
	}

	return &asset, nil
}

func (r *assetsRepository) GetAssetTicker(ctx context.Context, asset *models.Asset) (*models.AssetTicker, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "assets.repository.GetAssetTicker")
	defer span.Finish()

	var at models.AssetTicker
	at.Asset = *asset

	if err := r.db.GetContext(ctx, &at, getAssetTicker, at.ID); err != nil {
		return nil, err
	}

	return &at, nil
}

func (r *assetsRepository) SaveAssetTicker(ctx context.Context, assetTicker *models.AssetTicker) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "assets.repository.SaveAssetTicker")
	defer span.Finish()

	_, err := r.db.NamedExecContext(ctx, insertAssetTicker, &assetTicker)
	return err
}
