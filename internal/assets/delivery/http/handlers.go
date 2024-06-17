package http

import (
	"github/francoggm/crypto-wallets/config"
	"github/francoggm/crypto-wallets/internal/assets/usecase"

	"github.com/gofiber/fiber/v3"
)

type AssetsHandlers struct {
	cfg *config.Config
	uc  *usecase.AssetsUseCase
}

func NewHandlers(cfg *config.Config, uc *usecase.AssetsUseCase) *AssetsHandlers {
	return &AssetsHandlers{
		cfg,
		uc,
	}
}

func (h AssetsHandlers) ListAllAssetsData(c fiber.Ctx) error {
	return nil
}

func (h AssetsHandlers) ListAssetData(c fiber.Ctx) error {
	return nil
}
