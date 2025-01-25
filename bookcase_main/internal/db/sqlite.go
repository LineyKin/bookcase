package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

const SQLITE_DRIVER = "sqlite"
const SQLITE_DBFILE = "my_lib.db"

type Sqlite struct {
	Connection
}

func newSqlite() Connection {
	return &Sqlite{}
}

func (sqlt *Sqlite) Init() (*sql.DB, error) {
	db, err := sql.Open(SQLITE_DRIVER, SQLITE_DBFILE)
	if err != nil {
		log.Fatal("can't open database:", err)
		return nil, fmt.Errorf("can't open database: %s", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("can't connect to database:", err)
		return nil, fmt.Errorf("can't connect to database: %s", err)
	}

	return sqlt.createTables(db)

}

func (sqlt *Sqlite) createTables(db *sql.DB) (*sql.DB, error) {
	err := sqlt.createTableAuthors(db)
	if err != nil {
		return nil, err
	}

	err = sqlt.createTableBook(db)
	if err != nil {
		return nil, err
	}

	err = sqlt.createTableLiteraryWork(db)
	if err != nil {
		return nil, err
	}

	err = sqlt.createTableLiteraryWorkAndAuthors(db)
	if err != nil {
		return nil, err
	}

	err = sqlt.createTableLiteraryWorkAndBook(db)
	if err != nil {
		return nil, err
	}

	err = sqlt.createTablePublishingHouse(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// связная таблица литературных произведений и физических книг
func (sqlt *Sqlite) createTableLiteraryWorkAndBook(s *sql.DB) error {
	q := `
	CREATE TABLE IF NOT EXISTS book_and_literary_work (
		literary_work_id INTEGER,
		book_id INTEGER
	);`

	return sqlt.createTable(s, q, "book_and_literary_work")
}

// связная таблица литературных произведений и авторов
func (sqlt *Sqlite) createTableLiteraryWorkAndAuthors(s *sql.DB) error {
	q := `
	CREATE TABLE IF NOT EXISTS author_and_literary_work (
		author_id INTEGER,
		literary_work_id INTEGER
	);`

	return sqlt.createTable(s, q, "author_and_literary_work")
}

// таблица (физических) книг
func (sqlt *Sqlite) createTableBook(s *sql.DB) error {
	q := `
	CREATE TABLE IF NOT EXISTS book (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		year_of_publication INTEGER,
		publishing_house_id INTEGER
	);`

	return sqlt.createTable(s, q, "book")
}

// таблица литературных произведений
func (sqlt *Sqlite) createTablePublishingHouse(s *sql.DB) error {
	q := `
	CREATE TABLE IF NOT EXISTS publishing_house (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(256) NOT NULL DEFAULT ""
	);`

	return sqlt.createTable(s, q, "publishing_house")
}

// таблица литературных произведений
func (sqlt *Sqlite) createTableLiteraryWork(s *sql.DB) error {
	q := `
	CREATE TABLE IF NOT EXISTS literary_work (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(256) NOT NULL DEFAULT ""
	);`

	return sqlt.createTable(s, q, "literary_work")
}

// таблица авторов
func (sqlt *Sqlite) createTableAuthors(s *sql.DB) error {
	q := `
	CREATE TABLE IF NOT EXISTS authors (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(256) NOT NULL DEFAULT "",
		father_name VARCHAR(256) NOT NULL DEFAULT "",
		last_name VARCHAR(256) NOT NULL DEFAULT ""
	);`

	return sqlt.createTable(s, q, "authors")
}

func (sqlt *Sqlite) createTable(s *sql.DB, query, tableName string) error {
	_, err := s.Exec(query)
	if err != nil {
		return fmt.Errorf("can't create table `%s`: %w", tableName, err)
	}

	return nil
}
