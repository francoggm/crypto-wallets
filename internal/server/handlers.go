package server

import (
	assetsHttp "github/francoggm/crypto-wallets/internal/assets/delivery/http"
	assetsRepository "github/francoggm/crypto-wallets/internal/assets/repository"
	assetsUseCase "github/francoggm/crypto-wallets/internal/assets/usecase"

	"github.com/gofiber/fiber/v3"
)

func (s *Server) MapHandlers(router fiber.Router) error {
	// Repositorys
	assetsRepo := assetsRepository.NewRepository(s.db)

	// UseCases
	assetsUC := assetsUseCase.NewUseCase(assetsRepo)

	// Handlers
	assetsHandlers := assetsHttp.NewHandlers(s.cfg, assetsUC)

	// Assets routes
	assetsGp := router.Group("/assets")
	assetsHttp.MapRoutes(assetsGp, assetsHandlers)

	return nil
}
