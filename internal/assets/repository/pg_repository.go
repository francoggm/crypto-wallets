package repository

import (
	"github/francoggm/crypto-wallets/internal/assets"

	"github.com/jmoiron/sqlx"
)

type AssetsRepository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) assets.Repository {
	return &AssetsRepository{
		db: db,
	}
}
