package eventsRoutes

import (
	eventsController "hetmo_prueba_tecnica/internal/Events/infrastructure/controller"
	"hetmo_prueba_tecnica/internal/shared/infrastructure/middleware"

	"github.com/gofiber/fiber/v2"
)

func Init(r fiber.Router) {
	var controller eventsController.Events
	controller.New()

	//r.Get("/events", controller.GetEvents) // query params state - date - dentro del controller
	r.Get("/events/", controller.GetEventById)
	r.Get("/events/published", controller.GetPublishedEvents) // solo eventos futuros
	r.Get("/events/completed", controller.GetCompleteEvents)  // solo eventos pasados

	//protected routes
	AdminRoutes := r.Group("/events/admin", middleware.AdminMiddleware)
	AdminRoutes.Post("/new", controller.CreateEvent)
	AdminRoutes.Patch("/edit/", controller.UpdateEvent)
	AdminRoutes.Delete("/delete/:id", controller.DeleteEvent)

}
