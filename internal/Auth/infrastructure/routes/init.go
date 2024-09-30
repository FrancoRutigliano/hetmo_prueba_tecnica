package authRoutes

import (
	authController "hetmo_prueba_tecnica/internal/Auth/infrastructure/controller"
	"hetmo_prueba_tecnica/internal/shared/infrastructure/middleware"

	"github.com/gofiber/fiber/v2"
)

func Init(r fiber.Router) {
	var controller authController.Auth
	controller.NewAuthController()

	r.Post("/auth/register", controller.Register)
	r.Post("/auth/login", controller.Login)

	userRoutes := r.Group("/auth/user", middleware.UserMiddleware)
	userRoutes.Post("/change_password", controller.ChangePassword)

}
