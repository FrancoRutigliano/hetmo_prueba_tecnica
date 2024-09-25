package data

import (
	"fmt"
	"hetmo_prueba_tecnica/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnection(config *config.Config) (*sqlx.DB, error) {
	connStr := fmt.Sprintf("user=%s dbname=%s host=%s password=%s sslmode=disable", config.Postgre_User, config.Postgre_Db, config.Postgre_Host, config.Postgre_Password)

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
