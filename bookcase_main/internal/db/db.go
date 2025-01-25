package db

import (
	"database/sql"
	"fmt"
)

type Connection interface {
	Init() (*sql.DB, error)
}

type AppDB struct {
	Connection *sql.DB
	Driver     string
}

func Init(driver string) (AppDB, error) {
	conn, err := factory(driver)
	if err != nil {
		return AppDB{}, err
	}

	return AppDB{Connection: conn, Driver: driver}, nil

}

func factory(driver string) (*sql.DB, error) {
	switch driver {
	case PG_DRIVER:
		return newPostgres().Init()
	case SQLITE_DRIVER:
		return newSqlite().Init()
	default:
		return nil, fmt.Errorf("wrong db driver")
	}
}
