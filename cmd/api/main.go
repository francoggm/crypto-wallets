package main

import (
	"github/francoggm/crypto-wallets/config"
	"github/francoggm/crypto-wallets/internal/server"
	"log"
)

func main() {
	cfg := config.NewConfig()

	server := server.NewServer(cfg, nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
