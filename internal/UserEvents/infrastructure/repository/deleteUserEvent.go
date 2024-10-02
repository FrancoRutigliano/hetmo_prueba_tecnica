package infraUserEventsRepository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func (u *UserEventsImpl) DeleteUserEvent(userId, eventId string, db *sqlx.DB) error {
	query := `
	DELETE FROM "public".user_event 
	WHERE 
		user_id = $1 
		AND 
		event_id = $2;`

	result, err := db.Exec(query, userId, eventId)
	if err != nil {
		return fmt.Errorf("error deleting user event: %s", err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking affected rows")
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no event found with the provided ID")
	}

	return nil
}
