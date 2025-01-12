package db

import (
	"bookcase_log/lib/env"
	"database/sql"
	"fmt"
	"log"
)

const PG_DRIVER = "postgres"

func InitPostgresDb() (*sql.DB, error) {
	pgInfo := fmt.Sprintf("host=db2 port=%s user=%s password=%s dbname=%s sslmode=disable",
		env.GetPgPort(),
		env.GetPgUser(),
		env.GetPgPassword(),
		env.GetPgDbName(),
	)

	db, err := sql.Open(PG_DRIVER, pgInfo)
	if err != nil {
		return nil, err
	}

	// проверка связи с БД
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Postgres successfully connected")

	return db, nil
}
