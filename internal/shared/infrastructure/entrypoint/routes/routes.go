package routes

import (
	userRoutes "hetmo_prueba_tecnica/internal/User/infrastructure/routes"

	"github.com/gofiber/fiber/v2"
)

func Init(a *fiber.App) {
	api := a.Group("api/v1/")

	userRoutes.Init(api)
}
