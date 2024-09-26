package authRoutes

import (
	authController "hetmo_prueba_tecnica/internal/Auth/infrastructure/controller"
	authUseCase "hetmo_prueba_tecnica/internal/Auth/pkg/useCases"

	"github.com/gofiber/fiber/v2"
)

func Init(r fiber.Router) {
	authImpl := authUseCase.NewAuthUseCase()
	controller := authController.NewAuthController(authImpl)
	r.Post("login", controller.Login)
}
