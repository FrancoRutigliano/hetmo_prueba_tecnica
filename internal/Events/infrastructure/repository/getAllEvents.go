package infraEventsRepository

import (
	eventsDto "hetmo_prueba_tecnica/internal/Events/pkg/domain/dto"
	"time"

	"github.com/jmoiron/sqlx"
)

func (e *EventsImpl) GetAllEvents(isPublished *bool, date *time.Time, db *sqlx.DB) ([]eventsDto.EventListDTO, error) {
	query := `
	SELECT  
		e.title, 
		e.short_description, 
		e.date, 
		u.name AS organizer_name,
		e.location, 
		e.is_published
	FROM 
		"public".events AS e
	INNER JOIN 
		users AS u ON e.organizer = u.id
	WHERE 
		($1 IS NULL OR e.is_published = $1) 
		AND ($2 IS NULL OR e.date = $2);
	`

	var events []eventsDto.EventListDTO
	err := db.Select(&events, query, isPublished, date)
	if err != nil {
		return nil, err
	}
	return events, nil
}
