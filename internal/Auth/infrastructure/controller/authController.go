package authController

import (
	authDto "hetmo_prueba_tecnica/internal/Auth/pkg/domain/dto"
	authUseCase "hetmo_prueba_tecnica/internal/Auth/pkg/useCases"
	"hetmo_prueba_tecnica/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type Auth struct {
	handler authUseCase.AuthImpl
}

func (a *Auth) NewAuthController() {
	a.handler.New()
}

func (a *Auth) Login(c *fiber.Ctx) error {
	// sacar el body de la request
	var payload authDto.AuthLoginRequest

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request payload"})
	}

	if err := utils.Validate.Struct(&payload); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": "Invalid entity"})
	}

	// base de datos

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"response": "todo ok", "details": "true"})
}

func (a *Auth) Register(c *fiber.Ctx) error {
	var payload authDto.AuthRegisterRequest

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request payload", "details": "false"})
	}

	if err := utils.Validate.Struct(&payload); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": "Invalid entity", "details": "false"})
	}

	response, err := a.handler.AuthCase.Register(payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"response": err, err.Error(): "false"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"response": response, "details": "true"})
}
