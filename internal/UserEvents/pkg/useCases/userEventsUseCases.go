package userEventsUseCases

import (
	infraUserEventsRepository "hetmo_prueba_tecnica/internal/UserEvents/infrastructure/repository"
	userEventsUseCasesImpl "hetmo_prueba_tecnica/internal/UserEvents/pkg/useCases/useCasesImpl"
	"hetmo_prueba_tecnica/internal/shared/infrastructure/data"
	"log"
)

type EventsImpl struct {
	UserEventsCase userEventsUseCasesImpl.IUserEventsUseCase
}

func (a *EventsImpl) New() {
	var repository infraUserEventsRepository.SqlxUserEventsRepository

	repository.New()

	db, err := data.GetConnection()
	if err != nil {
		log.Println("error to connect db  --> ", err)
	}

	a.UserEventsCase = &userEventsUseCasesImpl.UserEvents{
		UserEvents: repository,
		Db:         db,
	}

}
