package eventsDto

import "time"

// EventCreateDTO es el DTO que se utiliza para crear un nuevo evento.
type EventCreateDTO struct {
	Title            string    `json:"title" validate:"required,max=100"`
	ShortDescription string    `json:"short_description" validate:"max=255"`
	LongDescription  string    `json:"long_description"`
	Date             time.Time `json:"date" validate:"required"`
	Organizer        string    `json:"organizer" validate:"required"`
	Location         string    `json:"location" validate:"required,max=100"`
	IsPublished      bool      `json:"is_published"`
}

// EventUpdateDTO es el DTO que se utiliza para actualizar un evento existente.
type EventUpdateDTO struct {
	Title            string    `json:"title,omitempty" validate:"max=100"`
	ShortDescription string    `json:"short_description,omitempty" validate:"max=255"`
	LongDescription  string    `json:"long_description,omitempty"`
	Date             time.Time `json:"date,omitempty"`
	Organizer        string    `json:"organizer,omitempty"`
	Location         string    `json:"location,omitempty" validate:"max=100"`
	IsPublished      bool      `json:"is_published,omitempty"`
}

// EventResponseDTO es el DTO que se utiliza para devolver la información de un evento específico.
type EventResponseDTO struct {
	ID               string    `json:"id"`                // ID único del evento.
	Title            string    `json:"title"`             // Título del evento.
	ShortDescription string    `json:"short_description"` // Descripción corta del evento.
	LongDescription  string    `json:"long_description"`  // Descripción larga del evento.
	Date             time.Time `json:"date"`              // Fecha y hora del evento.
	Organizer        string    `json:"organizer"`         // ID del organizador.
	Location         string    `json:"location"`          // Ubicación del evento.
	IsPublished      bool      `json:"is_published"`      // Indica si el evento está publicado o en borrador.
	CreatedAt        time.Time `json:"created_at"`        // Fecha de creación del evento.
	UpdatedAt        time.Time `json:"updated_at"`        // Fecha de última actualización.
}

// EventListDTO es el DTO que se utiliza para listar eventos, incluye información básica.
type EventListDTO struct {
	ID          string    `json:"id"`           // ID único del evento.
	Title       string    `json:"title"`        // Título del evento.
	Date        time.Time `json:"date"`         // Fecha y hora del evento.
	Location    string    `json:"location"`     // Ubicación del evento.
	IsPublished bool      `json:"is_published"` // Indica si el evento está publicado o no.
}
