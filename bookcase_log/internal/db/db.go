package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
)

const PG_DRIVER = "postgres"
const DUMP = "internal/db/log_table.sql"

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

	if !checkPgTables(db) {
		log.Print("need tables")
		file, err := os.ReadFile(DUMP)

		if err != nil {
			log.Fatal("can't read dump: ", err)
		}

		requests := strings.Split(string(file), ";")

		for _, request := range requests {
			_, err := db.Exec(request)
			if err != nil {
				log.Fatal("can't execute dump ", err)
			}
		}

	}

	return db, nil
}

func checkPgTables(db *sql.DB) bool {

	row := db.QueryRow("SELECT COUNT(tablename) AS count FROM pg_tables WHERE schemaname='public';")

	var count int

	_ = row.Scan(&count)

	return count > 0
}
