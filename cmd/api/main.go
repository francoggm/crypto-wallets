package main

import (
	"github/francoggm/crypto-wallets/config"
	"github/francoggm/crypto-wallets/internal/server"
	"github/francoggm/crypto-wallets/pkg/db/postgres"
	"github/francoggm/crypto-wallets/pkg/tickers"
	"log"
)

func main() {
	cfg := config.NewConfig()

	db, err := postgres.NewPostgres(cfg)
	if err != nil {
		log.Fatal(err)
	}

	tr := tickers.NewTickersRoutine(cfg, db)
	go tr.GetTickersData()

	server := server.NewServer(cfg, db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
