package middlewares

import (
	"github/francoggm/crypto-wallets/internal/auth"
	"github/francoggm/crypto-wallets/internal/models"
	"github/francoggm/crypto-wallets/pkg/utils"
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func (mw *MiddlewareManager) AuthMiddleware(c fiber.Ctx) error {
	token := c.Cookies("token")
	if token == "" {
		return c.Status(http.StatusUnauthorized).JSON(utils.GetMessageError(auth.ErrUserNotLogged))
	}

	claims, err := utils.ValidateTokenJWT(token, mw.cfg.Token.SecretKey)
	if err != nil {
		log.Error(err)
		return c.Status(http.StatusUnauthorized).JSON(utils.GetMessageError(auth.ErrUserNotLogged))
	}

	user, err := mw.repo.FindUserByID(c.Context(), claims.UserID)
	if err != nil {
		log.Error(err)
		return c.Status(http.StatusUnauthorized).JSON(utils.GetMessageError(auth.ErrUserNotLogged))
	}

	c.Locals("user", user)
	return c.Next()
}

func (mw *MiddlewareManager) AdminMiddleware(c fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok {
		return c.Status(http.StatusUnauthorized).JSON(utils.GetMessageError(auth.ErrInvalidPermission))
	}

	if !user.IsAdmin() {
		return c.Status(http.StatusUnauthorized).JSON(utils.GetMessageError(auth.ErrInvalidPermission))
	}

	return c.Next()
}
