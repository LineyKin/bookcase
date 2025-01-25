package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

const PG_DRIVER = "postgres"
const DUMP = "internal/db/bookcase_dump.sql"

type Postgres struct {
	Connection
}

func newPostgres() Connection {
	return &Postgres{}
}

func (pg *Postgres) Init() (*sql.DB, error) {
	pgInfo := fmt.Sprintf("host=db1 port=5432 user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_DBNAME"),
	)

	log.Println("pgInfo: ", pgInfo)

	db, err := sql.Open(PG_DRIVER, pgInfo)
	if err != nil {
		log.Print("ошибка подключения: ", err)
		return nil, err
	}

	// проверка связи с БД
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Postgres successfully connected")

	if !pg.checkPgTables(db) {
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

func (pg *Postgres) checkPgTables(db *sql.DB) bool {

	row := db.QueryRow("SELECT COUNT(tablename) AS count FROM pg_tables WHERE schemaname='public';")

	var count int

	_ = row.Scan(&count)

	return count > 0
}
