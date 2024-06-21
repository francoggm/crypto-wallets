package http

import (
	"github/francoggm/crypto-wallets/config"
	"github/francoggm/crypto-wallets/internal/auth"
	"github/francoggm/crypto-wallets/internal/models"
	"github/francoggm/crypto-wallets/pkg/utils"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/opentracing/opentracing-go"
)

type authHandlers struct {
	cfg *config.Config
	uc  auth.UseCase
}

func NewHandlers(cfg *config.Config, uc auth.UseCase) auth.Handlers {
	return &authHandlers{
		cfg: cfg,
		uc:  uc,
	}
}

func (h *authHandlers) Register() fiber.Handler {
	return func(c fiber.Ctx) error {
		span, ctx := opentracing.StartSpanFromContext(c.Context(), "handlers.auth.Register")
		defer span.Finish()

		user := new(models.User)
		if err := c.Bind().Body(user); err != nil {
			log.Error(err)
			return c.Status(http.StatusBadRequest).JSON(utils.GetMessageError(auth.ErrInvalidBody))
		}

		if !user.ValidUsername() {
			return c.Status(http.StatusBadRequest).JSON(utils.GetMessageError(auth.ErrInvalidUsername))
		}

		if !user.ValidEmail() {
			return c.Status(http.StatusBadRequest).JSON(utils.GetMessageError(auth.ErrInvalidEmail))
		}

		if !user.ValidPassword() {
			return c.Status(http.StatusBadRequest).JSON(utils.GetMessageError(auth.ErrInvalidPassword))
		}

		if err := h.uc.Register(ctx, user); err != nil {
			return c.Status(http.StatusBadRequest).JSON(utils.GetMessageError(err))
		}

		token, err := utils.GenerateTokenJWT(h.cfg.Token.Expiration, h.cfg.Token.SecretKey, user)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(utils.GetMessageError(auth.ErrFailedInLogin))
		}

		cookie := &fiber.Cookie{
			Name:     "token",
			Value:    token,
			Expires:  time.Now().Add(time.Duration(h.cfg.Token.Expiration) * time.Hour),
			HTTPOnly: true,
			Secure:   true,
		}
		c.Cookie(cookie)

		return c.JSON(utils.GetMessage("Registered!"))
	}
}

func (h *authHandlers) Login() fiber.Handler {
	return func(c fiber.Ctx) error {
		span, ctx := opentracing.StartSpanFromContext(c.Context(), "handlers.auth.Login")
		defer span.Finish()

		ul := new(models.UserLogin)
		if err := c.Bind().Body(ul); err != nil {
			return c.Status(http.StatusBadRequest).JSON(utils.GetMessageError(auth.ErrInvalidBody))
		}

		if !ul.ValidEmail() {
			return c.Status(http.StatusBadRequest).JSON(utils.GetMessageError(auth.ErrInvalidCredentials))
		}

		if !ul.ValidPassword() {
			return c.Status(http.StatusBadRequest).JSON(utils.GetMessageError(auth.ErrInvalidCredentials))
		}

		user, err := h.uc.Login(ctx, ul)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(utils.GetMessageError(err))
		}

		token, err := utils.GenerateTokenJWT(h.cfg.Token.Expiration, h.cfg.Token.SecretKey, user)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(utils.GetMessageError(auth.ErrFailedInLogin))
		}

		if err := h.uc.UpdateLastLogin(ctx, user.ID, time.Now()); err != nil {
			log.Error(err)
		}

		cookie := &fiber.Cookie{
			Name:     "token",
			Value:    token,
			Expires:  time.Now().Add(time.Duration(h.cfg.Token.Expiration) * time.Hour),
			HTTPOnly: true,
			Secure:   true,
		}
		c.Cookie(cookie)

		return c.JSON(utils.GetMessage("Logged in!"))
	}
}

func (h *authHandlers) Logout() fiber.Handler {
	return func(c fiber.Ctx) error {
		span, _ := opentracing.StartSpanFromContext(c.Context(), "handlers.auth.Logout")
		defer span.Finish()

		cookie := &fiber.Cookie{
			Name:     "token",
			Value:    "",
			Expires:  time.Now().Add(-time.Hour),
			HTTPOnly: true,
			Secure:   true,
		}
		c.Cookie(cookie)

		return c.JSON(utils.GetMessage("Logged out!"))
	}
}

func (h *authHandlers) Me() fiber.Handler {
	return func(c fiber.Ctx) error {
		span, _ := opentracing.StartSpanFromContext(c.Context(), "handlers.auth.Me")
		defer span.Finish()

		user, ok := c.Locals("user").(*models.User)
		if !ok {
			return c.Status(http.StatusUnauthorized).JSON(utils.GetMessageError(auth.ErrFailedGettingUserInfos))
		}

		user.Password = ""
		return c.JSON(user)
	}
}
