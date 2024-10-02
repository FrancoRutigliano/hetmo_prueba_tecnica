package infraEventsRepository

import (
	"fmt"
	eventsDto "hetmo_prueba_tecnica/internal/Events/pkg/domain/dto"

	"github.com/jmoiron/sqlx"
)

func (e *EventsImpl) GetEventById(id string, db *sqlx.DB) (eventsDto.EventListDTO, error) {
	query := `
	SELECT 
		e.title, 
		e.short_description, 
		to_timestamp(e.date) AS date,
		u.name AS organizer, 
		e.location, 
		e.is_published 
	FROM 
		"public".events AS e 
	INNER JOIN 
		users AS u ON e.organizer = u.id
	WHERE 
		e.id=$1`

	var payload eventsDto.EventListDTO

	err := db.Get(&payload, query, id)
	if err != nil {
		return eventsDto.EventListDTO{}, fmt.Errorf("error executing the query: %s", err.Error())
	}
	return payload, nil
}
