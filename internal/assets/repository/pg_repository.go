package repository

import (
	"github/francoggm/crypto-wallets/internal/assets"
	"github/francoggm/crypto-wallets/internal/models"
	"strings"

	"github.com/jmoiron/sqlx"
)

type assetsRepository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) assets.Repository {
	return &assetsRepository{
		db: db,
	}
}

func (r *assetsRepository) GetAllAssets() ([]*models.Asset, error) {
	var assets []*models.Asset

	if err := r.db.Select(&assets, getAllAssets); err != nil {
		return nil, err
	}

	return assets, nil
}

func (r *assetsRepository) GetAsset(assetName string) (*models.Asset, error) {
	var asset models.Asset

	if err := r.db.Get(&asset, getAsset, assetName, strings.ToLower(assetName), strings.ToUpper(assetName)); err != nil {
		return nil, err
	}

	return &asset, nil
}

func (r *assetsRepository) GetAssetTicker(asset models.Asset) (*models.AssetTicker, error) {
	var at models.AssetTicker
	at.Asset = asset

	if err := r.db.Get(&at, getAssetTicker, at.Id); err != nil {
		return nil, err
	}

	return &at, nil
}
