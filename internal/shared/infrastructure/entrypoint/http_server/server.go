package server

import (
	"hetmo_prueba_tecnica/config"
	"hetmo_prueba_tecnica/internal/shared/infrastructure/entrypoint/http_server/routes"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (s *Server) Run() error {
	app := fiber.New()

	routes.Init(app)

	return app.Listen(s.config.Port)
}
