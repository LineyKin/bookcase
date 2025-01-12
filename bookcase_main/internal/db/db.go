package db

import (
	"bookcase/lib/env"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

const PG_DRIVER = "postgres"
const SQLITE_DRIVER = "sqlite"
const DUMP = "internal/db/bookcase_dump.sql"

func InitPostgresDb() (*sql.DB, error) {
	pgInfo := fmt.Sprintf("host=db1 port=5432 user=%s password=%s dbname=%s sslmode=disable",
		env.GetPgUser(),
		env.GetPgPassword(),
		env.GetPgDbName(),
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

func InitSqliteDB() (*sql.DB, error) {
	db, err := sql.Open(SQLITE_DRIVER, env.GetDbName())
	if err != nil {
		log.Fatal("can't open database:", err)
		return nil, fmt.Errorf("can't open database: %s", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("can't connect to database:", err)
		return nil, fmt.Errorf("can't connect to database: %s", err)
	}

	return createTables(db)

}

func createTables(db *sql.DB) (*sql.DB, error) {
	err := createTableAuthors(db)
	if err != nil {
		return nil, err
	}

	err = createTableBook(db)
	if err != nil {
		return nil, err
	}

	err = createTableLiteraryWork(db)
	if err != nil {
		return nil, err
	}

	err = createTableLiteraryWorkAndAuthors(db)
	if err != nil {
		return nil, err
	}

	err = createTableLiteraryWorkAndBook(db)
	if err != nil {
		return nil, err
	}

	err = createTablePublishingHouse(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// связная таблица литературных произведений и физических книг
func createTableLiteraryWorkAndBook(s *sql.DB) error {
	q := `
	CREATE TABLE IF NOT EXISTS book_and_literary_work (
		literary_work_id INTEGER,
		book_id INTEGER
	);`

	return createTable(s, q, "book_and_literary_work")
}

// связная таблица литературных произведений и авторов
func createTableLiteraryWorkAndAuthors(s *sql.DB) error {
	q := `
	CREATE TABLE IF NOT EXISTS author_and_literary_work (
		author_id INTEGER,
		literary_work_id INTEGER
	);`

	return createTable(s, q, "author_and_literary_work")
}

// таблица (физических) книг
func createTableBook(s *sql.DB) error {
	q := `
	CREATE TABLE IF NOT EXISTS book (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		year_of_publication INTEGER,
		publishing_house_id INTEGER
	);`

	return createTable(s, q, "book")
}

// таблица литературных произведений
func createTablePublishingHouse(s *sql.DB) error {
	q := `
	CREATE TABLE IF NOT EXISTS publishing_house (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(256) NOT NULL DEFAULT ""
	);`

	return createTable(s, q, "publishing_house")
}

// таблица литературных произведений
func createTableLiteraryWork(s *sql.DB) error {
	q := `
	CREATE TABLE IF NOT EXISTS literary_work (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(256) NOT NULL DEFAULT ""
	);`

	return createTable(s, q, "literary_work")
}

// таблица авторов
func createTableAuthors(s *sql.DB) error {
	q := `
	CREATE TABLE IF NOT EXISTS authors (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(256) NOT NULL DEFAULT "",
		father_name VARCHAR(256) NOT NULL DEFAULT "",
		last_name VARCHAR(256) NOT NULL DEFAULT ""
	);`

	return createTable(s, q, "authors")
}

func createTable(s *sql.DB, query, tableName string) error {
	_, err := s.Exec(query)
	if err != nil {
		return fmt.Errorf("can't create table `%s`: %w", tableName, err)
	}

	return nil
}
