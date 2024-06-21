package server

import (
	"github/francoggm/crypto-wallets/config"

	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	cfg *config.Config
	app *fiber.App
	db  *sqlx.DB
}

func NewServer(cfg *config.Config, db *sqlx.DB) *Server {
	return &Server{
		cfg,
		fiber.New(),
		db,
	}
}

func (s *Server) Run() error {
	gp := s.app.Group("/v1")
	s.MapHandlers(gp)

	return s.app.Listen(":" + s.cfg.Server.Port)
}
