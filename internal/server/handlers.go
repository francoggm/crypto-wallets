package server

import (
	assetsHTTP "github/francoggm/crypto-wallets/internal/assets/delivery/http"
	assetsRepository "github/francoggm/crypto-wallets/internal/assets/repository"
	assetsUseCase "github/francoggm/crypto-wallets/internal/assets/usecase"
	authHTTP "github/francoggm/crypto-wallets/internal/auth/delivery/http"
	authRepository "github/francoggm/crypto-wallets/internal/auth/repository"
	authUseCase "github/francoggm/crypto-wallets/internal/auth/usecase"
	"github/francoggm/crypto-wallets/internal/middlewares"
	walletsHTTP "github/francoggm/crypto-wallets/internal/wallets/delivery/http"
	walletsRepository "github/francoggm/crypto-wallets/internal/wallets/repository"
	walletsUseCase "github/francoggm/crypto-wallets/internal/wallets/usecase"

	"github.com/gofiber/fiber/v3"
)

func (s *Server) MapHandlers(router fiber.Router) error {
	// Repositorys
	authRepo := authRepository.NewRepository(s.db)
	assetsRepo := assetsRepository.NewRepository(s.db)
	walletsRepo := walletsRepository.NewRepository(s.db)

	// UseCases
	authUC := authUseCase.NewUseCase(authRepo)
	assetsUC := assetsUseCase.NewUseCase(assetsRepo)
	walletsUC := walletsUseCase.NewUseCase(walletsRepo, assetsRepo)

	// Handlers
	authHandlers := authHTTP.NewHandlers(s.cfg, authUC)
	assetsHandlers := assetsHTTP.NewHandlers(s.cfg, assetsUC)
	walletsHandlers := walletsHTTP.NewHandlers(s.cfg, walletsUC)

	// Middlewares
	mw := middlewares.NewMiddlewareManager(s.cfg, authRepo)

	// Auth routes
	authGp := router.Group("/auth")
	authHTTP.MapRoutes(authGp, authHandlers, mw)

	// Assets routes
	assetsGp := router.Group("/assets")
	assetsHTTP.MapRoutes(assetsGp, assetsHandlers, mw)

	// Wallets routes
	walletsGp := router.Group("/wallets")
	walletsHTTP.MapRoutes(walletsGp, walletsHandlers, mw)

	return nil
}
