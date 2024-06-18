package http

import (
	"github/francoggm/crypto-wallets/config"
	"github/francoggm/crypto-wallets/internal/assets"

	"github.com/gofiber/fiber/v3"
)

type AssetsHandlers struct {
	cfg *config.Config
	uc  assets.UseCase
}

func NewHandlers(cfg *config.Config, uc assets.UseCase) assets.Handlers {
	return &AssetsHandlers{
		cfg,
		uc,
	}
}

func (h AssetsHandlers) ListAllAssetsData() fiber.Handler {
	return func(c fiber.Ctx) error {
		return nil
	}
}

func (h AssetsHandlers) ListAssetData() fiber.Handler {
	return func(c fiber.Ctx) error {
		return nil
	}
}
