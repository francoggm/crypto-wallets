package assets

import (
	"github.com/gofiber/fiber/v3"
)

type Handlers interface {
	GetAllAssetsData() fiber.Handler
	GetAssetData() fiber.Handler
	GetAssetHistoricalData() fiber.Handler
}
