package infraUserEventsRepository

import "github.com/jmoiron/sqlx"

func (u *UserEventsImpl) DeleteUserEvent(userId, eventId string, db *sqlx.DB) error {
	return nil
}
