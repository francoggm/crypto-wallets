package http

import (
	"github/francoggm/crypto-wallets/internal/middlewares"
	"github/francoggm/crypto-wallets/internal/wallets"

	"github.com/gofiber/fiber/v3"
)

func MapRoutes(gp fiber.Router, handlers wallets.Handlers, mw *middlewares.MiddlewareManager) {
	gp.Get("/", handlers.ListAllWallets(), mw.AuthMiddleware)
	gp.Get("/:asset", handlers.ListWallet(), mw.AuthMiddleware)
	gp.Post("/buy", handlers.BuyAsset(), mw.AuthMiddleware)
	gp.Post("/sell", handlers.SellAsset(), mw.AuthMiddleware)
}
