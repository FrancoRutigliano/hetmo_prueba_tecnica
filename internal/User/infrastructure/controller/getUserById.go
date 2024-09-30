package userController

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (u *User) GetUserById(c *fiber.Ctx) error {

	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "user id required", "details": "false"})
	}

	response := u.handler.UserCase.FindUserById(id)
	if response.StatusCode != http.StatusOK {
		return c.Status(response.StatusCode).JSON(fiber.Map{"message": response.Msg, "details": "false"})
	}
	return c.Status(response.StatusCode).JSON(fiber.Map{"message": response.Msg, "data": response.Data, "details": "true"})
}
