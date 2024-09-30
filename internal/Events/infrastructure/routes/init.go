package eventsRoutes

import (
	eventsController "hetmo_prueba_tecnica/internal/Events/infrastructure/controller"

	"github.com/gofiber/fiber/v2"
)

func Init(r fiber.Router) {
	var controller eventsController.Events
	controller.New()

	r.Post("/events/new", controller.CreateEvent)
	r.Get("/events", controller.GetEvents) // query params state - date - dentro del controller
	r.Get("/events/:id", controller.GetEventById)
	r.Put("/events/edit/:id", controller.UpdateEvent)
	r.Delete("/events/delete/:id", controller.DeleteEvent)
	r.Get("/events/published", controller.GetPublishedEvents) // solo eventos futuros
	r.Get("/events/completed", controller.GetCompleteEvents)  // solo eventos pasados

}
