package userRoutes

import (
	userController "hetmo_prueba_tecnica/internal/User/infrastructure/controller"

	"github.com/gofiber/fiber/v2"
)

func Init(r fiber.Router) {
	var controller userController.User
	controller.NewUserController()

	r.Get("/users/:id", controller.GetUserById)
	r.Put("/users/edit/:id", controller.UpdateUser)
	r.Delete("/users/delete/:id", controller.DeleteUser)
}
