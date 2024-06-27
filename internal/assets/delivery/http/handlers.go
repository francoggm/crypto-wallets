package http

import (
	"database/sql"
	"errors"
	"github/francoggm/crypto-wallets/config"
	"github/francoggm/crypto-wallets/internal/assets"
	"github/francoggm/crypto-wallets/pkg/utils"
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/opentracing/opentracing-go"
)

type assetsHandlers struct {
	cfg *config.Config
	uc  assets.UseCase
}

func NewHandlers(cfg *config.Config, uc assets.UseCase) assets.Handlers {
	return &assetsHandlers{
		cfg,
		uc,
	}
}

func (h assetsHandlers) ListAllAssetsData() fiber.Handler {
	return func(c fiber.Ctx) error {
		span, ctx := opentracing.StartSpanFromContext(c.Context(), "assets.handlers.ListAllAssetsData")
		defer span.Finish()

		assetsTickers, err := h.uc.ListAllAssetsData(ctx)
		if err != nil {
			log.Error(err)
			return c.Status(http.StatusBadRequest).Send(nil)
		}

		if len(assetsTickers) == 0 {
			return c.Status(http.StatusBadRequest).JSON(utils.GetMessageError(assets.ErrFailedGettingTicker))
		}

		return c.JSON(assetsTickers)
	}
}

func (h assetsHandlers) ListAssetData() fiber.Handler {
	return func(c fiber.Ctx) error {
		span, ctx := opentracing.StartSpanFromContext(c.Context(), "assets.handlers.ListAssetData")
		defer span.Finish()

		assetName := c.Params("asset")
		if assetName == "" {
			return c.Status(http.StatusUnprocessableEntity).JSON(utils.GetMessageError(assets.ErrInvalidAsset))
		}

		assetTicker, err := h.uc.ListAssetData(ctx, assetName)
		if err != nil {
			log.Error(err)

			if errors.Is(err, sql.ErrNoRows) {
				return c.Status(http.StatusNotFound).Send(nil)
			}

			return c.Status(http.StatusBadRequest).JSON(utils.GetMessageError(assets.ErrFailedGettingTicker))
		}

		return c.JSON(assetTicker)
	}
}
