package eventsDto

import "time"

// EventCreateDTO es el DTO que se utiliza para crear un nuevo evento.
type EventCreateDTO struct {
	Title            string    `json:"title" validate:"required,max=100"`    // El título del evento es obligatorio y con longitud máxima de 100 caracteres.
	ShortDescription string    `json:"short_description" validate:"max=255"` // Descripción corta opcional con longitud máxima de 255 caracteres.
	LongDescription  string    `json:"long_description"`                     // Descripción larga opcional.
	Date             time.Time `json:"date" validate:"required"`             // La fecha es obligatoria.
	Organizer        string    `json:"organizer" validate:"required"`
	Location         string    `json:"location" validate:"required,max=100"` // La ubicación es obligatoria con longitud máxima de 100 caracteres.
	IsPublished      bool      `json:"is_published"`                         // Estado del evento (borrador o publicado).
}

// EventUpdateDTO es el DTO que se utiliza para actualizar un evento existente.
type EventUpdateDTO struct {
	Title            string    `json:"title,omitempty" validate:"max=100"`             // Permite actualizar opcionalmente el título.
	ShortDescription string    `json:"short_description,omitempty" validate:"max=255"` // Permite actualizar opcionalmente la descripción corta.
	LongDescription  string    `json:"long_description,omitempty"`                     // Permite actualizar opcionalmente la descripción larga.
	Date             time.Time `json:"date,omitempty"`                                 // Permite actualizar opcionalmente la fecha.
	Organizer        string    `json:"organizer,omitempty"`
	Location         string    `json:"location,omitempty" validate:"max=100"` // Permite actualizar opcionalmente la ubicación.
	IsPublished      bool      `json:"is_published,omitempty"`                // Permite actualizar opcionalmente el estado.
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
