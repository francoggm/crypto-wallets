package assets

import "github.com/gofiber/fiber/v3"

type Handlers interface {
	ListAllAssetsData() fiber.Handler
	ListAssetData() fiber.Handler
}
