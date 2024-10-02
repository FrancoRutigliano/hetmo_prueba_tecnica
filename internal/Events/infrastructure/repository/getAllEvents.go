package infraEventsRepository

import (
	eventsDto "hetmo_prueba_tecnica/internal/Events/pkg/domain/dto"

	"github.com/jmoiron/sqlx"
)

func (e *EventsImpl) GetAllEvents(db *sqlx.DB) ([]eventsDto.EventListDTO, error) {
	// Definir la consulta para obtener todos los eventos
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
		users AS u ON e.organizer = u.id;`

	// Ejecutar la consulta
	var events []eventsDto.EventListDTO
	err := db.Select(&events, query)
	if err != nil {
		return nil, err
	}

	return events, nil
}
