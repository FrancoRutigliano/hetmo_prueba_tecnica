package userRoutes

import "github.com/gofiber/fiber/v2"

func Init(r fiber.Router) {
	r.Get("/login", func(c *fiber.Ctx) error {
		return c.SendString("Login")
	})
}
