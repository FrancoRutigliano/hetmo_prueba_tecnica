package infraEventsRepository

import (
	"fmt"
	eventsDto "hetmo_prueba_tecnica/internal/Events/pkg/domain/dto"
	eventsRepository "hetmo_prueba_tecnica/internal/Events/pkg/domain/repository"
	"time"

	"github.com/jmoiron/sqlx"
)

type SqlxEventsRepository struct {
	Impl eventsRepository.Repository
}

type EventsImpl struct {
}

func (s *SqlxEventsRepository) New() {
	s.Impl = &EventsImpl{}
}

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

// terminar
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
func (e *EventsImpl) GetEventById(id string, db *sqlx.DB) (eventsDto.EventListDTO, error) {
	query := `
	SELECT 
		e.title, 
		e.short_description, 
		e.date,
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

	// Ejecuta la consulta y devuelve el evento actualizado
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

func (e *EventsImpl) DeleteEvent(id string, db *sqlx.DB) error {
	query := `DELETE FROM "public".events WHERE id=$1`
	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no event found with the provided ID")
	}

	return nil
}

func (e *EventsImpl) GetCompletedEvents(db *sqlx.DB) ([]eventsDto.EventListDTO, error) {
	return []eventsDto.EventListDTO{}, nil
}

func (e *EventsImpl) GetPublishedEvents(db *sqlx.DB) ([]eventsDto.EventListDTO, error) {
	return []eventsDto.EventListDTO{}, nil
}
