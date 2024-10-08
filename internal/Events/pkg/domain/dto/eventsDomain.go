package eventsDto

import "time"

// EventCreateDTO es el DTO que se utiliza para crear un nuevo evento.
type EventCreateDTO struct {
	Id               string    `json:"id"`
	Title            string    `json:"title" `
	ShortDescription string    `json:"short_description"`
	LongDescription  string    `json:"long_description"`
	Date             int64     `json:"date" `
	Organizer        string    `json:"organizer"`
	Location         string    `json:"location"`
	IsPublished      bool      `json:"is_published"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type EventCreateDTORequest struct {
	Title            string `json:"title" validate:"required,max=100"`
	ShortDescription string `json:"short_description" validate:"max=255"`
	LongDescription  string `json:"long_description"`
	Date             string `json:"date" validate:"required"`
	Organizer        string `json:"organizer"`
	Location         string `json:"location" validate:"required,max=100"`
	IsPublished      bool   `json:"is_published"`
}

type GetEventsRequest struct {
	Title       string
	IsPublished bool
	Date        int64
}

// EventResponseDTO es el DTO que se utiliza para devolver la información de un evento específico.
type EventResponseDTO struct {
	Title            string    `db:"title" json:"title,omitempty" validate:"max=100"`
	ShortDescription string    `db:"short_description" json:"short_description,omitempty" validate:"max=255"`
	LongDescription  string    `db:"long_description" json:"long_description,omitempty"`
	Date             string    `db:"date" json:"date,omitempty"`
	Organizer        string    `db:"organizer" json:"organizer"`
	Location         string    `db:"location" json:"location,omitempty" validate:"max=100"`
	IsPublished      bool      `db:"is_published" json:"is_published,omitempty"`
	UpdatedAt        time.Time `db:"updated_at" json:"updated_at"`
}

type EventResponseDTOUpdate struct {
	Title            string    `db:"title" json:"title,omitempty" validate:"max=100"`
	ShortDescription string    `db:"short_description" json:"short_description,omitempty" validate:"max=255"`
	LongDescription  string    `db:"long_description" json:"long_description,omitempty"`
	Date             int64     `db:"date" json:"date,omitempty"`
	Organizer        string    `db:"organizer" json:"organizer"`
	Location         string    `db:"location" json:"location,omitempty" validate:"max=100"`
	IsPublished      bool      `db:"is_published" json:"is_published,omitempty"`
	UpdatedAt        time.Time `db:"updated_at" json:"updated_at"`
}

// EventListDTO es el DTO que se utiliza para listar eventos, incluye información básica.
type EventListDTO struct {
	Title            string    `db:"title" json:"title"`
	ShortDescription string    `db:"short_description" json:"short_description"`
	Date             time.Time `db:"date" json:"date"`
	Organizer        string    `db:"organizer" json:"organizer"`
	Location         string    `db:"location" json:"location"`
	IsPublished      bool      `db:"is_published" json:"is_published"`
}

type EventListDTOUpdate struct {
	Title            string `db:"title" json:"title"`
	ShortDescription string `db:"short_description" json:"short_description"`
	Date             int64  `db:"date" json:"date"`
	Organizer        string `db:"organizer" json:"organizer"`
	Location         string `db:"location" json:"location"`
	IsPublished      bool   `db:"is_published" json:"is_published"`
}
