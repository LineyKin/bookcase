package storage

import (
	"bookcase_log/internal/storage/db/postgres"
	"bookcase_log/models"
	"database/sql"
	"time"
)

type StorageInterface interface {
	AddLog(lr models.LogRow) error
	GetLatestLogTimestamp() time.Time
}

type Storage struct {
	StorageInterface
}

func New(db *sql.DB) *Storage {
	return &Storage{
		StorageInterface: postgres.New(db),
	}
}
