package infraAuthRepository

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func (s *ImplrepositoryAuth) FindByEmail(emailParam string, db *sqlx.DB) error {
	query := `SELECT u.email FROM "public".users AS u WHERE email=$1 LIMIT 1;`

	var email string

	if err := db.Get(&email, query, emailParam); err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return fmt.Errorf("error executing the query")
	}
	return fmt.Errorf("email: %s already exist", emailParam)
}
