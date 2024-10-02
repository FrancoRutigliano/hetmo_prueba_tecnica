package userEventsController

import userEventsUseCases "hetmo_prueba_tecnica/internal/UserEvents/pkg/useCases"

type UserEvents struct {
	handler userEventsUseCases.EventsImpl
}

func (e *UserEvents) New() {
	e.handler.New()
}
