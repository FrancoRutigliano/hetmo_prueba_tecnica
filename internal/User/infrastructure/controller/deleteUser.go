package userController

import "github.com/gofiber/fiber/v2"

func (u *User) DeleteUser(c *fiber.Ctx) error {
	return c.Status(201).JSON(fiber.Map{"message": "delete users"})
}
