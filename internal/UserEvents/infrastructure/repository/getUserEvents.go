package infraUserEventsRepository

import (
	userEventsDto "hetmo_prueba_tecnica/internal/UserEvents/pkg/domain/dto"

	"github.com/jmoiron/sqlx"
)

func (u *UserEventsImpl) GetUserEvent(userId, eventId string, db *sqlx.DB) ([]userEventsDto.UserEventsListDto, error) {
	return []userEventsDto.UserEventsListDto{}, nil
}
