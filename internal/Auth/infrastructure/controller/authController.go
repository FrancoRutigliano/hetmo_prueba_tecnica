package authController

import (
	authDomain "hetmo_prueba_tecnica/internal/Auth/pkg/domain"
	authUseCase "hetmo_prueba_tecnica/internal/Auth/pkg/useCases"
	"hetmo_prueba_tecnica/pkg/utils"

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
	// sacar el body de la request
	var payload authDomain.AuthLoginRequest

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	if err := utils.Validate.Struct(&payload); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": "Invalid entity"})
	}

	a.auth.Login(payload)

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"response": "todo ok", "details": "true"})
}

func (a *authController) Register(c *fiber.Ctx) error {

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"response": "user", "details": "true"})
}
