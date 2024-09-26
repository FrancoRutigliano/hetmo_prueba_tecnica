package server

import (
	"hetmo_prueba_tecnica/config"
	"hetmo_prueba_tecnica/internal/shared/infrastructure/entrypoint/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	config *config.Config
	db     *sqlx.DB
}

func NewServer(config *config.Config, db *sqlx.DB) *Server {
	return &Server{
		config: config,
		db:     db,
	}
}

func (s *Server) Run() error {
	app := fiber.New()

	routes.Init(app)

	return app.Listen(s.config.Port)
}
