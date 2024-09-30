package routes

import (
	authRoutes "hetmo_prueba_tecnica/internal/Auth/infrastructure/routes"
	eventsRoutes "hetmo_prueba_tecnica/internal/Events/infrastructure/routes"
	userRoutes "hetmo_prueba_tecnica/internal/User/infrastructure/routes"
	userEventsRoutes "hetmo_prueba_tecnica/internal/UserEvents/infrastructure/routes"

	"github.com/gofiber/fiber/v2"
)

func Init(a *fiber.App) {
	api := a.Group("api/v1/")
	authRoutes.Init(api)
	userRoutes.Init(api)
	eventsRoutes.Init(api)
	userEventsRoutes.Init(api)
}
