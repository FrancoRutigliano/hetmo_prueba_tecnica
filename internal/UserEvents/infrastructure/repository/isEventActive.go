package infraUserEventsRepository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func (u *UserEventsImpl) IsEventActive(eventID string, db *sqlx.DB) (bool, error) {
	// Consulta SQL para verificar si el evento está activo
	query := `
		SELECT
			COUNT(*) > 0
		FROM
			"public".events AS e
		WHERE
			e.id = $1
			AND e.is_published = TRUE
			AND e.date > EXTRACT(EPOCH FROM NOW())::BIGINT;
	`

	// Ejecutar la consulta
	var isActive bool
	err := db.Get(&isActive, query, eventID)
	if err != nil {
		return false, fmt.Errorf("error checking event active status: %v", err)
	}

	return isActive, nil
}
