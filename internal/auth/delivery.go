package auth

import "github.com/gofiber/fiber/v3"

type Handlers interface {
	Register() fiber.Handler
	Login() fiber.Handler
	Logout() fiber.Handler
	Me() fiber.Handler
}
