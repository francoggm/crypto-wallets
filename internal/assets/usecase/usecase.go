package usecase

import "github/francoggm/crypto-wallets/internal/assets/repository"

type AssetsUseCase struct {
	repo *repository.AssetsRepository
}

func NewUseCase(repo *repository.AssetsRepository) *AssetsUseCase {
	return &AssetsUseCase{
		repo,
	}
}
