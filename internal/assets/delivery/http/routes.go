package http

import (
	"github/francoggm/crypto-wallets/internal/assets"
	"github/francoggm/crypto-wallets/internal/middlewares"

	"github.com/gofiber/fiber/v3"
)

func MapRoutes(gp fiber.Router, handlers assets.Handlers, mw *middlewares.MiddlewareManager) {
	gp.Get("/", handlers.GetAllAssetsData())
	gp.Get("/:asset", handlers.GetAssetData())
	gp.Get("/:asset/history", handlers.GetAssetHistoricalData())
}
