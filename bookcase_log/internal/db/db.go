package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

const PG_DRIVER = "postgres"

func InitPostgresDb() (*sql.DB, error) {
	pgInfo := fmt.Sprintf("host=db2 port=5432 user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_DBNAME_LOG"),
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
