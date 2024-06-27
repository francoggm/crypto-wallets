package usecase

import (
	"context"
	"database/sql"
	"errors"
	"github/francoggm/crypto-wallets/internal/assets"
	"github/francoggm/crypto-wallets/internal/models"
	"github/francoggm/crypto-wallets/internal/wallets"

	"github.com/gofiber/fiber/v3/log"
	"github.com/opentracing/opentracing-go"
)

const (
	buyTransaction  = "buy"
	sellTransaction = "sell"
)

type walletsUseCase struct {
	walletsRepo wallets.Repository
	assetsRepo  assets.Repository
}

func NewUseCase(walletsRepo wallets.Repository, assetsRepo assets.Repository) wallets.UseCase {
	return &walletsUseCase{
		walletsRepo: walletsRepo,
		assetsRepo:  assetsRepo,
	}
}

func (uc *walletsUseCase) GetUserWallets(ctx context.Context, userId int) ([]*models.WalletTicker, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "wallets.usecase.GetUserWallets")
	defer span.Finish()

	wallets, err := uc.walletsRepo.GetUserWallets(ctx, userId)
	if err != nil {
		return nil, err
	}

	walletsTicker := make([]*models.WalletTicker, 0, len(wallets))

	for _, wallet := range wallets {
		asset, err := uc.assetsRepo.GetAssetByID(ctx, wallet.AssetID)
		if err != nil {
			log.Error(err)
			continue
		}

		assetTicker, err := uc.assetsRepo.GetAssetTicker(ctx, asset)
		if err != nil {
			log.Error(err)
			continue
		}

		walletTicker := models.WalletTicker{
			Symbol:    asset.Symbol,
			Asset:     asset.Name,
			PriceUSD:  assetTicker.PriceUSD,
			Quantity:  wallet.Quantity,
			TotalUSD:  assetTicker.PriceUSD * wallet.Quantity,
			UpdatedAt: wallet.UpdatedAt,
		}

		walletsTicker = append(walletsTicker, &walletTicker)
	}

	return walletsTicker, nil
}

func (uc *walletsUseCase) GetUserWallet(ctx context.Context, userId int, assetName string) (*models.WalletTicker, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "wallets.usecase.GetUserWallet")
	defer span.Finish()

	asset, err := uc.assetsRepo.GetAssetByName(ctx, assetName)
	if err != nil {
		return nil, err
	}

	wallet, err := uc.walletsRepo.GetUserWallet(ctx, userId, asset.ID)
	if err != nil {
		return nil, err
	}

	assetTicker, err := uc.assetsRepo.GetAssetTicker(ctx, asset)
	if err != nil {
		return nil, err
	}

	walletTicker := models.WalletTicker{
		Symbol:    asset.Symbol,
		Asset:     asset.Name,
		PriceUSD:  assetTicker.PriceUSD,
		Quantity:  wallet.Quantity,
		TotalUSD:  assetTicker.PriceUSD * wallet.Quantity,
		UpdatedAt: wallet.UpdatedAt,
	}

	return &walletTicker, nil
}

func (uc *walletsUseCase) BuyAsset(ctx context.Context, user *models.User, transaction *models.TransactionAsset) (*models.WalletTicker, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "wallets.usecase.BuyAsset")
	defer span.Finish()

	asset, err := uc.assetsRepo.GetAssetByName(ctx, transaction.Asset)
	if err != nil {
		return nil, err
	}

	assetTicker, err := uc.assetsRepo.GetAssetTicker(ctx, asset)
	if err != nil {
		return nil, err
	}

	price := transaction.Quantity * assetTicker.PriceUSD

	if price > user.WalletUSD {
		return nil, wallets.ErrInsufficientAmount
	}

	wallet, err := uc.walletsRepo.GetUserWallet(ctx, user.ID, asset.ID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	if wallet == nil {
		wallet, err = uc.walletsRepo.CreateWallet(ctx, user.ID, asset.ID)
		if err != nil {
			return nil, err
		}
	}

	wallet, err = uc.walletsRepo.UpdateWallet(ctx, wallet.Quantity+transaction.Quantity, wallet)
	if err != nil {
		return nil, err
	}

	err = uc.walletsRepo.UpdateWalletUSD(ctx, user.ID, user.WalletUSD-price)
	if err != nil {
		log.Error(err)
	}

	err = uc.walletsRepo.CreateTransaction(ctx, user.ID, asset.ID, buyTransaction, transaction.Quantity, assetTicker.PriceUSD)
	if err != nil {
		log.Error(err)
	}

	walletTicker := models.WalletTicker{
		Symbol:    asset.Symbol,
		Asset:     asset.Name,
		PriceUSD:  assetTicker.PriceUSD,
		Quantity:  transaction.Quantity,
		TotalUSD:  price,
		UpdatedAt: wallet.UpdatedAt,
	}

	return &walletTicker, nil
}

func (uc *walletsUseCase) SellAsset(ctx context.Context, user *models.User, transaction *models.TransactionAsset) (*models.WalletTicker, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "wallets.usecase.SellAsset")
	defer span.Finish()

	asset, err := uc.assetsRepo.GetAssetByName(ctx, transaction.Asset)
	if err != nil {
		return nil, err
	}

	wallet, err := uc.walletsRepo.GetUserWallet(ctx, user.ID, asset.ID)
	if err != nil {
		return nil, err
	}

	if wallet.Quantity < transaction.Quantity {
		return nil, wallets.ErrInsufficientAmount
	}

	wallet, err = uc.walletsRepo.UpdateWallet(ctx, wallet.Quantity-transaction.Quantity, wallet)
	if err != nil {
		return nil, err
	}

	assetTicker, err := uc.assetsRepo.GetAssetTicker(ctx, asset)
	if err != nil {
		return nil, err
	}

	price := transaction.Quantity * assetTicker.PriceUSD

	err = uc.walletsRepo.UpdateWalletUSD(ctx, user.ID, user.WalletUSD+price)
	if err != nil {
		log.Error(err)
	}

	err = uc.walletsRepo.CreateTransaction(ctx, user.ID, asset.ID, sellTransaction, transaction.Quantity, assetTicker.PriceUSD)
	if err != nil {
		log.Error(err)
	}

	walletTicker := models.WalletTicker{
		Symbol:    asset.Symbol,
		Asset:     asset.Name,
		PriceUSD:  assetTicker.PriceUSD,
		Quantity:  transaction.Quantity,
		TotalUSD:  price,
		UpdatedAt: wallet.UpdatedAt,
	}

	return &walletTicker, nil
}
