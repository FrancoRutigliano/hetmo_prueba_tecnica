package infraUserEventsRepository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func (u *UserEventsImpl) CreateUserEvent(userId, eventId string, db *sqlx.DB) error {
	// Definir la consulta SQL de inserción
	query := `
		INSERT INTO user_event (user_id, event_id, is_active, created_at)
		VALUES ($1, $2, $3, NOW())
		ON CONFLICT (user_id, event_id) DO NOTHING;
	`

	_, err := db.Exec(query, userId, eventId, true)
	if err != nil {
		return fmt.Errorf("failed to create user-event relation: %v", err)
	}

	return nil
}
