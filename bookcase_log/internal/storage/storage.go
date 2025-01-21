package storage

import (
	"bookcase_log/internal/storage/db/postgres"
	"bookcase_log/models"
	"database/sql"
)

type StorageInterface interface {
	AddLog(lr models.LogRow) error
}

type Storage struct {
	StorageInterface
}

func New(db *sql.DB) *Storage {
	return &Storage{
		StorageInterface: postgres.New(db),
	}
}
