package infraEventsRepository

import (
	eventsDto "hetmo_prueba_tecnica/internal/Events/pkg/domain/dto"

	"github.com/jmoiron/sqlx"
)

func (e *EventsImpl) GetPublishedEvents(title string, db *sqlx.DB) ([]eventsDto.EventListDTO, error) {
	var args []interface{}
	var query string

	// Consulta base
	query = `
	SELECT
		e.title,
		e.short_description,
		e.date,
		u.name AS organizer,
		e.location,
		e.is_published
	FROM
		"public".events AS e
	LEFT JOIN
		users AS u ON e.organizer = u.id
	WHERE
		e.is_published = TRUE
		AND e.date > EXTRACT(EPOCH FROM NOW())::BIGINT
	`

	if title != "" {
		query += ` AND e.title ILIKE '%' || $1 || '%';`
		args = append(args, title) // Agrega el título a los argumentos
	}

	var payload []eventsDto.EventListDTO
	err := db.Select(&payload, query, args...)
	if err != nil {
		return nil, err
	}
	return payload, nil
}
