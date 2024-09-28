package authRoutes

import (
	authController "hetmo_prueba_tecnica/internal/Auth/infrastructure/controller"

	"github.com/gofiber/fiber/v2"
)

func Init(r fiber.Router) {
	var controller authController.Auth
	controller.NewAuthController()

	r.Post("/register", controller.Register)
	r.Post("login", controller.Login)
}
