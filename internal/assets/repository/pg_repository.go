package repository

import "github.com/jmoiron/sqlx"

type AssetsRepository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *AssetsRepository {
	return &AssetsRepository{
		db: db,
	}
}
