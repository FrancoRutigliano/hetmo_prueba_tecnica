package infraEventsRepository

import (
	"fmt"
	eventsDto "hetmo_prueba_tecnica/internal/Events/pkg/domain/dto"

	"github.com/jmoiron/sqlx"
)

func (e *EventsImpl) UpdateEvent(eventId string, updatedEvent eventsDto.EventResponseDTO, db *sqlx.DB) (eventsDto.EventResponseDTO, error) {
	// Consulta para actualizar el evento
	updateQuery := `
		UPDATE "public".events 
		SET 
			title = $1, 
			short_description = $2, 
			long_description = $3, 
			date = $4, 
			organizer = $5, 
			location = $6, 
			is_published = $7, 
			updated_at = NOW()
		WHERE 
			id = $8
		RETURNING title, short_description, long_description, date, organizer, location, is_published, updated_at
	`

	var modifiedEvent eventsDto.EventResponseDTO
	err := db.Get(&modifiedEvent, updateQuery,
		updatedEvent.Title,
		updatedEvent.ShortDescription,
		updatedEvent.LongDescription,
		updatedEvent.Date,
		updatedEvent.Organizer,
		updatedEvent.Location,
		updatedEvent.IsPublished,
		eventId,
	)

	if err != nil {
		return eventsDto.EventResponseDTO{}, fmt.Errorf("error updating the event: %s", err.Error())
	}

	return modifiedEvent, nil
}
