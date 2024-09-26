package authController

import (
	authUseCase "hetmo_prueba_tecnica/internal/Auth/pkg/useCases"

	"github.com/gofiber/fiber/v2"
)

type authController struct {
	auth authUseCase.AuthUseCase
}

func NewAuthController(auth *authUseCase.AuthUseCaseImpl) *authController {
	return &authController{
		auth: auth,
	}
}

func (a *authController) Login(c *fiber.Ctx) error {
	response, err := a.auth.Login()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "internal"})
	}

	return c.Status(200).JSON(fiber.Map{"response": response})
}
