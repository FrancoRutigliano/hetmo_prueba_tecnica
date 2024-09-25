package server

import (
	"hetmo_prueba_tecnica/config"

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

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	return app.Listen(s.config.Port)
}
