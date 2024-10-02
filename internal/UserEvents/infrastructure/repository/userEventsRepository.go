package infraUserEventsRepository

import (
	userEventsRepository "hetmo_prueba_tecnica/internal/UserEvents/pkg/domain/repository"
)

type SqlxUserEventsRepository struct {
	Impl userEventsRepository.Repository
}

type UserEventsImpl struct {
}

func (s *SqlxUserEventsRepository) New() {
	s.Impl = &UserEventsImpl{}
}
