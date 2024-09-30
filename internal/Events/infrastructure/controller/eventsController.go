package eventsController

import eventsUseCase "hetmo_prueba_tecnica/internal/Events/pkg/useCases"

type Events struct {
	handler eventsUseCase.EventsImpl
}

func (e *Events) New() {
	e.handler.New()
}
