package infraEventsRepository

import (
	"fmt"
	eventsDto "hetmo_prueba_tecnica/internal/Events/pkg/domain/dto"

	"github.com/jmoiron/sqlx"
)

func (e *EventsImpl) GetAllEvents(title string, isPublished bool, date int64, db *sqlx.DB) ([]eventsDto.EventListDTO, error) {
	// Definir la consulta base
	query := `
	SELECT  
		e.title, 
		e.short_description, 
		to_timestamp(e.date),
		u.name AS organizer_name,
		e.location, 
		e.is_published
	FROM 
		"public".events AS e
	INNER JOIN 
		users AS u ON e.organizer = u.id
	WHERE 1=1` // Usar `1=1` para simplificar el agregado de condiciones dinámicas

	var args []interface{}
	argIndex := 1 // Índice para los placeholders

	// Filtrar por título si se especifica
	if title != "" {
		query += fmt.Sprintf(" AND e.title ILIKE $%d", argIndex)
		args = append(args, "%"+title+"%") // Agregar el título con wildcards para ILIKE
		argIndex++
	}

	// Filtrar por estado de publicación (isPublished)
	query += fmt.Sprintf(" AND e.is_published = %d", argIndex)
	args = append(args, isPublished)
	argIndex++

	// Filtrar por fecha si se especifica
	if date != 0 {
		query += fmt.Sprint(" AND DATE(TO_TIMESTAMP(e.date)) = DATE(TO_TIMESTAMP(EXTRACT(EPOCH FROM NOW())::BIGINT))", argIndex)
		args = append(args, date) // Pasar la fecha directamente
		argIndex++
	}

	query += " ORDER BY e.date DESC;" // Ordenar los eventos por fecha descendente

	// Ejecutar la consulta
	var events []eventsDto.EventListDTO
	err := db.Select(&events, query, args...)
	if err != nil {
		return nil, err
	}

	return events, nil
}
