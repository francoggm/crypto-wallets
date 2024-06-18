package usecase

import (
	"github/francoggm/crypto-wallets/internal/assets"
)

type AssetsUseCase struct {
	repo assets.Repository
}

func NewUseCase(repo assets.Repository) assets.UseCase {
	return &AssetsUseCase{
		repo,
	}
}
