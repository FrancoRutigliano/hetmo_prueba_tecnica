package infraEventsRepository

import (
	eventsDto "hetmo_prueba_tecnica/internal/Events/pkg/domain/dto"

	"github.com/jmoiron/sqlx"
)

func (e *EventsImpl) CreateEvent(payload eventsDto.EventCreateDTO, db *sqlx.DB) error {
	query := `INSERT INTO  "public".events (id, title, short_description, long_description, date, organizer, location, is_published, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);`

	_, err := db.Exec(query, payload.Id, payload.Title, payload.ShortDescription, payload.LongDescription,
		payload.Date, payload.Organizer, payload.Location, payload.IsPublished, payload.CreatedAt, payload.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}
