package eventsUseCase

import (
	infraEventsRepository "hetmo_prueba_tecnica/internal/Events/infrastructure/repository"
	eventsUseCaseImpl "hetmo_prueba_tecnica/internal/Events/pkg/useCases/useCaseImpl"
	"hetmo_prueba_tecnica/internal/shared/infrastructure/data"
	"log"
)

type EventsImpl struct {
	EventsCase eventsUseCaseImpl.IEventsUseCase
}

func (a *EventsImpl) New() {
	var repository infraEventsRepository.SqlxEventsRepository

	repository.New()

	db, err := data.GetConnection()
	if err != nil {
		log.Println("error to connect db  --> ", err)
	}

	a.EventsCase = &eventsUseCaseImpl.Events{
		Repository: repository,
		Db:         db,
	}

}
