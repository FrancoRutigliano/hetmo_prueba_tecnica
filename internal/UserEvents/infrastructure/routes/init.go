package userEventsRoutes

import (
	userEventsController "hetmo_prueba_tecnica/internal/UserEvents/infrastructure/controller"
	"hetmo_prueba_tecnica/internal/shared/infrastructure/middleware"

	"github.com/gofiber/fiber/v2"
)

func Init(r fiber.Router) {
	var controller = userEventsController.UserEvents{}

	controller.New()

	userRoutes := r.Group("/user_events/user", middleware.UserMiddleware)
	userRoutes.Get("/", controller.GeteUserEvent)
	userRoutes.Post("/new", controller.CreateUserEvent)
	userRoutes.Delete("/delete", controller.DeleteUserEvent)

}
