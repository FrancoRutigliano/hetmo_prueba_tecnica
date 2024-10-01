package infraEventsRepository

import (
	eventsRepository "hetmo_prueba_tecnica/internal/Events/pkg/domain/repository"
)

type SqlxEventsRepository struct {
	Impl eventsRepository.Repository
}

type EventsImpl struct {
}

func (s *SqlxEventsRepository) New() {
	s.Impl = &EventsImpl{}
}
