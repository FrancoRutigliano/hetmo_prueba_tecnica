package eventsUseCase

import (
	infraEventsRepository "hetmo_prueba_tecnica/internal/Events/infrastructure/repository"
	eventsUseCaseImpl "hetmo_prueba_tecnica/internal/Events/pkg/useCases/useCaseImpl"
	infraUserRepository "hetmo_prueba_tecnica/internal/User/infrastructure/repository"
	"hetmo_prueba_tecnica/internal/shared/infrastructure/data"
	"log"
)

type EventsImpl struct {
	EventsCase eventsUseCaseImpl.IEventsUseCase
}

func (a *EventsImpl) New() {
	var eventsRepository infraEventsRepository.SqlxEventsRepository
	eventsRepository.New()
	var userRepository infraUserRepository.SqlxUserRepository
	userRepository.New()
	db, err := data.GetConnection()
	if err != nil {
		log.Println("error to connect db  --> ", err)
	}

	a.EventsCase = &eventsUseCaseImpl.Events{
		EventsRepository: eventsRepository,
		Db:               db,
	}

}
