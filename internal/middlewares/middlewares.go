package middlewares

import (
	"github/francoggm/crypto-wallets/config"
	"github/francoggm/crypto-wallets/internal/auth"
)

type MiddlewareManager struct {
	cfg  *config.Config
	repo auth.Repository
}

func NewMiddlewareManager(cfg *config.Config, repo auth.Repository) *MiddlewareManager {
	return &MiddlewareManager{
		cfg:  cfg,
		repo: repo,
	}
}
