package main

import (
	"github/francoggm/crypto-wallets/config"
	"github/francoggm/crypto-wallets/internal/server"
	"github/francoggm/crypto-wallets/pkg/db/postgres"
	"log"
)

func main() {
	cfg := config.NewConfig()

	db, err := postgres.NewPostgres(cfg)
	if err != nil {
		log.Fatal(err)
	}

	server := server.NewServer(cfg, db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
