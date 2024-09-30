package eventsRepository

import (
	eventsDto "hetmo_prueba_tecnica/internal/Events/pkg/domain/dto"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	CreateEvent(eventsDto.EventCreateDTO, *sqlx.DB) error    // recibe dto
	GetAllEvents(eventsDto.EventResponseDTO, *sqlx.DB) error //  reciba objeto con los filtros
	GetEventById(string, *sqlx.DB) error                     // dto única devolución
	UpdateEvent(string, *sqlx.DB) error
	DeleteEvent(string, *sqlx.DB) error
	GetPublishedEvents(*sqlx.DB) ([]eventsDto.EventListDTO, error) //dto []arreglo
	GetCompletedEvents(*sqlx.DB) ([]eventsDto.EventListDTO, error) //dto []
}
