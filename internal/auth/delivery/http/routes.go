package http

import (
	"github/francoggm/crypto-wallets/internal/auth"
	"github/francoggm/crypto-wallets/internal/middlewares"

	"github.com/gofiber/fiber/v3"
)

func MapRoutes(gp fiber.Router, handlers auth.Handlers, mw *middlewares.MiddlewareManager) {
	gp.Post("/register", handlers.Register())
	gp.Post("/login", handlers.Login())
	gp.Post("/logout", handlers.Logout(), mw.AuthMiddleware)
	gp.Get("/me", handlers.Me(), mw.AuthMiddleware)
}
