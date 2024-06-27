package wallets

import "github.com/gofiber/fiber/v3"

type Handlers interface {
	ListAllWallets() fiber.Handler
	ListWallet() fiber.Handler
	BuyAsset() fiber.Handler
	SellAsset() fiber.Handler
}
