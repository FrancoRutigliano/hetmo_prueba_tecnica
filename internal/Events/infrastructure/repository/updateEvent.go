package infraEventsRepository

import (
	"fmt"
	eventsDto "hetmo_prueba_tecnica/internal/Events/pkg/domain/dto"

	"github.com/jmoiron/sqlx"
)

func (e *EventsImpl) UpdateEvent(eventId string, updatedEvent eventsDto.EventResponseDTOUpdate, db *sqlx.DB) (eventsDto.EventListDTO, error) {
	// Consulta para actualizar el evento
	updateQuery := `
		UPDATE "public".events 
		SET 
			title = $1, 
			short_description = $2, 
			date = $3, 
			organizer = $4, 
			location = $5, 
			is_published = $6, 
			updated_at = NOW()
		WHERE 
			id = $7
		RETURNING title, short_description, date, organizer, location, is_published
	`

	var modifiedEvent eventsDto.EventListDTO
	err := db.Get(&modifiedEvent, updateQuery,
		updatedEvent.Title,
		updatedEvent.ShortDescription,
		updatedEvent.Date,
		updatedEvent.Organizer,
		updatedEvent.Location,
		updatedEvent.IsPublished,
		eventId,
	)

	if err != nil {
		return eventsDto.EventListDTO{}, fmt.Errorf("error updating the event: %s", err.Error())
	}

	return modifiedEvent, nil
}
