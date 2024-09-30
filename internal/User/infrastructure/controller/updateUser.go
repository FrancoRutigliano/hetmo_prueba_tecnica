package userController

import "github.com/gofiber/fiber/v2"

func (u *User) UpdateUser(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{"message": "update users"})
}
