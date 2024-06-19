package http

import (
	"database/sql"
	"errors"
	"github/francoggm/crypto-wallets/config"
	"github/francoggm/crypto-wallets/internal/assets"
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
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
		assetsTickers, err := h.uc.ListAllAssetsData()
		if err != nil {
			log.Error(err)
			return c.Status(http.StatusBadRequest).Send(nil)
		}

		if len(assetsTickers) == 0 {
			return c.Status(http.StatusBadRequest).JSON(map[string]string{
				"message": "Error getting tickers, try again!",
			})
		}

		return c.JSON(assetsTickers)
	}
}

func (h assetsHandlers) ListAssetData() fiber.Handler {
	return func(c fiber.Ctx) error {
		assetName := c.Params("asset")
		if assetName == "" {
			return c.Status(http.StatusUnprocessableEntity).JSON(map[string]string{
				"message": "Invalid asset",
			})
		}

		assetTicker, err := h.uc.ListAssetData(assetName)
		if err != nil {
			log.Error(err)

			if errors.Is(err, sql.ErrNoRows) {
				return c.Status(http.StatusNotFound).Send(nil)
			}

			return c.Status(http.StatusBadRequest).JSON(map[string]string{
				"message": "Error getting ticker, try again!",
			})
		}

		return c.JSON(assetTicker)
	}
}
