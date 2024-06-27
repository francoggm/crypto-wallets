package http

import (
	"database/sql"
	"errors"
	"github/francoggm/crypto-wallets/config"
	"github/francoggm/crypto-wallets/internal/assets"
	"github/francoggm/crypto-wallets/internal/auth"
	"github/francoggm/crypto-wallets/internal/models"
	"github/francoggm/crypto-wallets/internal/wallets"
	"github/francoggm/crypto-wallets/pkg/utils"
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/opentracing/opentracing-go"
)

type walletsHandlers struct {
	cfg *config.Config
	uc  wallets.UseCase
}

func NewHandlers(cfg *config.Config, uc wallets.UseCase) wallets.Handlers {
	return &walletsHandlers{
		cfg: cfg,
		uc:  uc,
	}
}

func (h *walletsHandlers) ListAllWallets() fiber.Handler {
	return func(c fiber.Ctx) error {
		span, ctx := opentracing.StartSpanFromContext(c.Context(), "wallets.handlers.GetUserWallets")
		defer span.Finish()

		user, ok := c.Locals("user").(*models.User)
		if !ok {
			return c.Status(http.StatusUnauthorized).JSON(utils.GetMessageError(auth.ErrUserNotLogged))
		}

		walletsTickers, err := h.uc.GetUserWallets(ctx, user.ID)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			log.Error(err)
			return c.Status(http.StatusBadRequest).JSON(utils.GetMessageError(wallets.ErrFailedGettingUserWallets))
		}

		return c.JSON(walletsTickers)
	}
}

func (h *walletsHandlers) ListWallet() fiber.Handler {
	return func(c fiber.Ctx) error {
		span, ctx := opentracing.StartSpanFromContext(c.Context(), "wallets.handlers.GetUserWallet")
		defer span.Finish()

		user, ok := c.Locals("user").(*models.User)
		if !ok {
			return c.Status(http.StatusUnauthorized).JSON(utils.GetMessageError(auth.ErrUserNotLogged))
		}

		asset := c.Params("asset")
		if asset == "" {
			return c.Status(http.StatusBadRequest).JSON(utils.GetMessageError(assets.ErrInvalidAsset))
		}

		walletTicker, err := h.uc.GetUserWallet(ctx, user.ID, asset)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return c.Status(http.StatusNotFound).JSON(utils.GetMessageError(wallets.ErrWalletNotFound))
			}

			log.Error(err)
			return c.Status(http.StatusBadRequest).JSON(utils.GetMessageError(wallets.ErrFailedGettingUserWallet))
		}

		return c.JSON(walletTicker)
	}
}

func (h *walletsHandlers) BuyAsset() fiber.Handler {
	return func(c fiber.Ctx) error {
		span, ctx := opentracing.StartSpanFromContext(c.Context(), "wallets.handlers.BuyAsset")
		defer span.Finish()

		user, ok := c.Locals("user").(*models.User)
		if !ok {
			return c.Status(http.StatusUnauthorized).JSON(utils.GetMessageError(auth.ErrUserNotLogged))
		}

		transaction := new(models.TransactionAsset)
		if err := c.Bind().Body(transaction); err != nil {
			return c.Status(http.StatusBadRequest).JSON(utils.GetMessageError(wallets.ErrInvalidBodyValues))
		}

		if transaction.Asset == "" {
			return c.Status(http.StatusUnprocessableEntity).JSON(utils.GetMessageError(assets.ErrInvalidAsset))
		}

		if transaction.Quantity <= 0 {
			return c.Status(http.StatusUnprocessableEntity).JSON(utils.GetMessageError(wallets.ErrInvalidAmount))
		}

		wt, err := h.uc.BuyAsset(ctx, user, transaction)
		if err != nil {
			if errors.Is(err, wallets.ErrInsufficientAmount) {
				return c.Status(http.StatusUnprocessableEntity).JSON(utils.GetMessageError(wallets.ErrInsufficientAmount))
			}

			log.Error(err)
			return c.Status(http.StatusBadRequest).JSON(utils.GetMessageError(wallets.ErrFailedBuyingAsset))
		}

		return c.JSON(wt)
	}
}

func (h *walletsHandlers) SellAsset() fiber.Handler {
	return func(c fiber.Ctx) error {
		span, ctx := opentracing.StartSpanFromContext(c.Context(), "wallets.handlers.SellAsset")
		defer span.Finish()

		user, ok := c.Locals("user").(*models.User)
		if !ok {
			return c.Status(http.StatusUnauthorized).JSON(utils.GetMessageError(auth.ErrUserNotLogged))
		}

		transaction := new(models.TransactionAsset)
		if err := c.Bind().Body(transaction); err != nil {
			return c.Status(http.StatusBadRequest).JSON(utils.GetMessageError(wallets.ErrInvalidBodyValues))
		}

		if transaction.Asset == "" {
			return c.Status(http.StatusUnprocessableEntity).JSON(utils.GetMessageError(assets.ErrInvalidAsset))
		}

		if transaction.Quantity <= 0 {
			return c.Status(http.StatusUnprocessableEntity).JSON(utils.GetMessageError(wallets.ErrInvalidAmount))
		}

		wt, err := h.uc.SellAsset(ctx, user, transaction)
		if err != nil {
			if errors.Is(err, wallets.ErrInsufficientAmount) {
				return c.Status(http.StatusUnprocessableEntity).JSON(utils.GetMessageError(wallets.ErrInsufficientAmount))
			}

			log.Error(err)
			return c.Status(http.StatusBadRequest).JSON(utils.GetMessageError(wallets.ErrFailedBuyingAsset))
		}

		return c.JSON(wt)
	}
}
