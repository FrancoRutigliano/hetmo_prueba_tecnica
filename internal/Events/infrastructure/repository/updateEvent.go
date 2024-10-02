package infraEventsRepository

import (
	"fmt"
	eventsDto "hetmo_prueba_tecnica/internal/Events/pkg/domain/dto"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

func (e *EventsImpl) UpdateEvent(eventId string, updatedEvent eventsDto.EventResponseDTOUpdate, db *sqlx.DB) (eventsDto.EventListDTOUpdate, error) {
	// Construir dinámicamente los campos que se actualizarán
	setClauses := []string{"organizer = $1"} // Incluir organizer como el primer campo siempre
	args := []interface{}{updatedEvent.Organizer}
	i := 2 // Comenzar desde el segundo índice ya que el organizer es $1

	// Solo agregar los campos que no están vacíos o con valores por defecto
	if updatedEvent.Title != "" {
		setClauses = append(setClauses, fmt.Sprintf("title = $%d", i))
		args = append(args, updatedEvent.Title)
		i++
	}
	if updatedEvent.ShortDescription != "" {
		setClauses = append(setClauses, fmt.Sprintf("short_description = $%d", i))
		args = append(args, updatedEvent.ShortDescription)
		i++
	}
	if updatedEvent.Date != 0 {
		setClauses = append(setClauses, fmt.Sprintf("date = $%d", i))
		args = append(args, updatedEvent.Date)
		i++
	}
	if updatedEvent.Location != "" {
		setClauses = append(setClauses, fmt.Sprintf("location = $%d", i))
		args = append(args, updatedEvent.Location)
		i++
	}
	if updatedEvent.IsPublished {
		setClauses = append(setClauses, fmt.Sprintf("is_published = $%d", i))
		args = append(args, updatedEvent.IsPublished)
		i++
	}

	// Actualizar el campo "updated_at" con la fecha actual
	setClauses = append(setClauses, fmt.Sprintf("updated_at = $%d", i))
	args = append(args, time.Now())
	i++

	// Agregar el ID del evento como último argumento
	args = append(args, eventId)

	// Construir la consulta de actualización dinámica
	updateQuery := fmt.Sprintf(`
		UPDATE "public".events
		SET %s
		WHERE id = $%d
		RETURNING title, short_description, date, organizer, location, is_published
	`, strings.Join(setClauses, ", "), i)

	// Definir la variable donde se almacenarán los resultados del RETURNING
	var modifiedEvent eventsDto.EventListDTOUpdate

	// Ejecutar la consulta y llenar modifiedEvent con los datos del evento actualizado
	err := db.Get(&modifiedEvent, updateQuery, args...)
	if err != nil {
		return eventsDto.EventListDTOUpdate{}, fmt.Errorf("error updating the event: %s", err.Error())
	}

	return modifiedEvent, nil
}
