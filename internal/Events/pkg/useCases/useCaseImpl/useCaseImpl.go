package eventsUseCaseImpl

import (
	infraEventsRepository "hetmo_prueba_tecnica/internal/Events/infrastructure/repository"
	eventsDto "hetmo_prueba_tecnica/internal/Events/pkg/domain/dto"
	httpresponse "hetmo_prueba_tecnica/pkg/httpResponse"

	"github.com/jmoiron/sqlx"
)

type IEventsUseCase interface {
	CreateEvent(eventsDto.EventCreateDTORequest, string, int64) httpresponse.ApiResponse
	GetEvents() httpresponse.ApiResponse
	GetEventByID(string) httpresponse.ApiResponse
	UpdateEvent(string, string, eventsDto.EventResponseDTO) httpresponse.ApiResponse // dto
	DeleteEvent(string) httpresponse.ApiResponse
	GetDraftEvents(string) httpresponse.ApiResponse
	GetCompletedEvents(string) httpresponse.ApiResponse
}

type Events struct {
	EventsRepository infraEventsRepository.SqlxEventsRepository
	Db               *sqlx.DB
}
