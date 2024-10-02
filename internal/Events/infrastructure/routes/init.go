package eventsRoutes

import (
	eventsController "hetmo_prueba_tecnica/internal/Events/infrastructure/controller"
	"hetmo_prueba_tecnica/internal/shared/infrastructure/middleware"

	"github.com/gofiber/fiber/v2"
)

func Init(r fiber.Router) {
	var controller eventsController.Events
	controller.New()

	r.Get("/events/all/", controller.GetEvents)
	r.Get("/events/", controller.GetEventById)
	r.Get("/events/completed", controller.GetCompleteEvents)

	//protected routes
	AdminRoutes := r.Group("/events/admin", middleware.AdminMiddleware)
	AdminRoutes.Get("/draft", controller.GetDraftEvents)
	AdminRoutes.Post("/new", controller.CreateEvent)
	AdminRoutes.Patch("/edit/", controller.UpdateEvent)
	AdminRoutes.Delete("/delete/", controller.DeleteEvent)

}
