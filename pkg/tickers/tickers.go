package tickers

import (
	"github/francoggm/crypto-wallets/config"
	"github/francoggm/crypto-wallets/internal/assets"
	"github/francoggm/crypto-wallets/internal/assets/repository"
	"time"

	"github.com/jmoiron/sqlx"
)

type TickersRoutine struct {
	TickersInterval time.Duration
	TickersURL      string
	Repository      assets.Repository
}

func NewTickersRoutine(cfg *config.Config, db *sqlx.DB) *TickersRoutine {
	repo := repository.NewRepository(db)

	return &TickersRoutine{
		TickersInterval: time.Duration(cfg.Tickers.Interval) * time.Second,
		TickersURL:      cfg.Tickers.URL,
		Repository:      repo,
	}
}
