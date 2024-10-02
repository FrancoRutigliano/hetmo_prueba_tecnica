package userEventsController

import "github.com/gofiber/fiber/v2"

func (u *UserEvents) GeteUserEvent(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{"message": "get"})
}
