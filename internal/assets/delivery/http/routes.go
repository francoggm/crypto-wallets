package http

import "github.com/gofiber/fiber/v3"

func MapRoutes(gp fiber.Router, handlers *AssetsHandlers) {
	gp.Get("/all", handlers.ListAllAssetsData)
	gp.Get("/:asset", handlers.ListAssetData)
}
