package repository

import (
	"context"
	"github/francoggm/crypto-wallets/internal/models"
	"github/francoggm/crypto-wallets/internal/wallets"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
)

type walletsRepository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) wallets.Repository {
	return &walletsRepository{
		db: db,
	}
}

func (r *walletsRepository) GetUserWallets(ctx context.Context, userId int) ([]*models.Wallet, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "wallets.repository.GetUserWallets")
	defer span.Finish()

	var wallets []*models.Wallet

	if err := r.db.SelectContext(ctx, &wallets, getUserWallets, userId); err != nil {
		return nil, err
	}

	return wallets, nil
}

func (r *walletsRepository) GetUserWallet(ctx context.Context, userId int, assetId int) (*models.Wallet, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "wallets.repository.GetUserWallet")
	defer span.Finish()

	var wallet models.Wallet

	if err := r.db.GetContext(ctx, &wallet, getUserWallet, userId, assetId); err != nil {
		return nil, err
	}

	return &wallet, nil
}

func (r *walletsRepository) CreateWallet(ctx context.Context, userId int, assetId int) (*models.Wallet, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "wallets.repository.CreateWallet")
	defer span.Finish()

	var wallet models.Wallet

	rows, err := r.db.NamedQueryContext(ctx, insertWallet, map[string]any{"user_id": userId, "asset_id": assetId})
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		err = rows.Scan(&wallet.ID, &wallet.UserID, &wallet.AssetID, &wallet.Quantity, &wallet.CreatedAt, &wallet.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &wallet, nil
}

func (r *walletsRepository) UpdateWallet(ctx context.Context, quantity float64, wallet *models.Wallet) (*models.Wallet, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "wallets.repository.UpdateWallet")
	defer span.Finish()

	var w models.Wallet

	rows, err := r.db.NamedQueryContext(ctx, updateWallet, map[string]any{"quantity": quantity, "user_id": wallet.UserID, "asset_id": wallet.AssetID, "updated_at": time.Now()})
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		err = rows.Scan(&w.ID, &w.UserID, &w.AssetID, &w.Quantity, &w.CreatedAt, &w.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &w, nil
}

func (r *walletsRepository) UpdateWalletUSD(ctx context.Context, userId int, amount float64) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "wallets.repository.UpdateUserWalletUSD")
	defer span.Finish()

	_, err := r.db.ExecContext(ctx, updateWalletUSD, userId, amount, time.Now())
	return err
}

func (r *walletsRepository) CreateTransaction(ctx context.Context, userId int, assetId int, transactionType string, amount float64, price float64) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "wallets.repository.CreateTransaction")
	defer span.Finish()

	_, err := r.db.ExecContext(ctx, insertTransaction, userId, assetId, transactionType, amount, price)
	return err
}
