package eventsUseCaseImpl

import (
	infraEventsRepository "hetmo_prueba_tecnica/internal/Events/infrastructure/repository"
	httpresponse "hetmo_prueba_tecnica/pkg/httpResponse"

	"github.com/jmoiron/sqlx"
)

type IEventsUseCase interface {
	CreateEvent() *httpresponse.ApiResponse // dto
	GetEvents() *httpresponse.ApiResponse   // dto
	GetEventByID(string) *httpresponse.ApiResponse
	UpdateEvent(string) *httpresponse.ApiResponse // dto
	DeleteEvent(string) *httpresponse.ApiResponse
	GetPublishedEvents() *httpresponse.ApiResponse
	GetCompletedEvents() *httpresponse.ApiResponse
}

type Events struct {
	Repository infraEventsRepository.SqlxEventsRepository
	Db         *sqlx.DB
}
