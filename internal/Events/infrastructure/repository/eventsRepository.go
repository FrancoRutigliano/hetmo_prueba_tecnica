package infraEventsRepository

import (
	eventsDto "hetmo_prueba_tecnica/internal/Events/pkg/domain/dto"
	eventsRepository "hetmo_prueba_tecnica/internal/Events/pkg/domain/repository"

	"github.com/jmoiron/sqlx"
)

type SqlxEventsRepository struct {
	Impl eventsRepository.Repository
}

type EventsImpl struct {
}

func (s *SqlxEventsRepository) New() {
	s.Impl = &EventsImpl{}
}

func (e *EventsImpl) CreateEvent(payload eventsDto.EventCreateDTO, db *sqlx.DB) error {
	return nil
}

func (e *EventsImpl) GetAllEvents(payload eventsDto.EventResponseDTO, db *sqlx.DB) error {
	return nil
}

func (e *EventsImpl) GetEventById(id string, db *sqlx.DB) error {
	return nil
}

func (e *EventsImpl) UpdateEvent(id string, db *sqlx.DB) error {
	return nil
}

func (e *EventsImpl) DeleteEvent(id string, db *sqlx.DB) error {
	return nil
}

func (e *EventsImpl) GetCompletedEvents(db *sqlx.DB) ([]eventsDto.EventListDTO, error) {
	return []eventsDto.EventListDTO{}, nil
}

func (e *EventsImpl) GetPublishedEvents(db *sqlx.DB) ([]eventsDto.EventListDTO, error) {
	return []eventsDto.EventListDTO{}, nil
}
