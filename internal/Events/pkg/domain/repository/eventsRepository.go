package eventsRepository

import (
	eventsDto "hetmo_prueba_tecnica/internal/Events/pkg/domain/dto"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	CreateEvent(eventsDto.EventCreateDTO, *sqlx.DB) error // recibe dto
	GetAllEvents(string, bool, int64, *sqlx.DB) ([]eventsDto.EventListDTO, error)
	GetEventById(string, *sqlx.DB) (eventsDto.EventListDTO, error) // dto única devolución
	UpdateEvent(string, eventsDto.EventResponseDTOUpdate, *sqlx.DB) (eventsDto.EventListDTOUpdate, error)
	DeleteEvent(string, *sqlx.DB) error
	GetPublishedEvents(string, *sqlx.DB) ([]eventsDto.EventListDTO, error) //dto []arreglo
	GetCompletedEvents(string, *sqlx.DB) ([]eventsDto.EventListDTO, error) //dto []
}
