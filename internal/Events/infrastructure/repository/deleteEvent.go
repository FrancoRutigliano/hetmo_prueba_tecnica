package infraEventsRepository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func (e *EventsImpl) DeleteEvent(id string, db *sqlx.DB) error {
	query := `DELETE FROM "public".events WHERE id=$1`
	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no event found with the provided ID")
	}

	return nil
}
