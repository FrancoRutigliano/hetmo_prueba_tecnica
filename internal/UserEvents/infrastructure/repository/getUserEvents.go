package infraUserEventsRepository

import (
	userEventsDto "hetmo_prueba_tecnica/internal/UserEvents/pkg/domain/dto"

	"github.com/jmoiron/sqlx"
)

func (u *UserEventsImpl) GetUserEvent(userId, eventId string, db *sqlx.DB) ([]userEventsDto.UserEventsListDto, error) {
	query := `
SELECT 
    e.id AS event_id,                -- Alias correcto
    e.title,
    e.short_description,
    DATE(TO_TIMESTAMP(e.date)) AS date,
    u.name AS organizer,
    e.location,
    e.is_published
FROM
    user_event AS ue
INNER JOIN 
    events AS e ON ue.event_id = e.id
INNER JOIN 
    users AS u ON e.organizer = u.id
WHERE
    ue.user_id = $1
    AND
    ue.is_active = TRUE;
`

	var events []userEventsDto.UserEventsListDto
	err := db.Select(&events, query, userId)
	if err != nil {
		return nil, err
	}

	return events, nil
}
