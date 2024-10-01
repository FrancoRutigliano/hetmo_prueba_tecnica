package infraEventsRepository

import (
	eventsDto "hetmo_prueba_tecnica/internal/Events/pkg/domain/dto"

	"github.com/jmoiron/sqlx"
)

func (e *EventsImpl) GetCompletedEvents(db *sqlx.DB) ([]eventsDto.EventListDTO, error) {
	return []eventsDto.EventListDTO{}, nil
}
