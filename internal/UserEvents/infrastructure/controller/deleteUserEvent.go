package userEventsController

import "github.com/gofiber/fiber/v2"

func (u *UserEvents) DeleteUserEvent(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{"message": "delete"})
}
