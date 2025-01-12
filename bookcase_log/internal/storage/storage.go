package storage

import (
	"bookcase_log/internal/storage/db/postgres"
	"database/sql"
)

type StorageInterface interface {
	AddLog() (int, error)
}

type Storage struct {
	StorageInterface
}

func New(db *sql.DB) *Storage {
	return &Storage{
		StorageInterface: postgres.New(db),
	}
}
