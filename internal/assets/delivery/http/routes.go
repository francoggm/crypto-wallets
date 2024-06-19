package http

import (
	"github/francoggm/crypto-wallets/internal/assets"

	"github.com/gofiber/fiber/v3"
)

func MapRoutes(gp fiber.Router, handlers assets.Handlers) {
	gp.Get("/", handlers.ListAllAssetsData())
	gp.Get("/:asset", handlers.ListAssetData())
}
