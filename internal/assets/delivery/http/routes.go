package http

import (
	"github/francoggm/crypto-wallets/internal/assets"
	"github/francoggm/crypto-wallets/internal/middlewares"

	"github.com/gofiber/fiber/v3"
)

func MapRoutes(gp fiber.Router, handlers assets.Handlers, mw *middlewares.MiddlewareManager) {
	gp.Get("/", handlers.ListAllAssetsData())
	gp.Get("/:asset", handlers.ListAssetData())
}
