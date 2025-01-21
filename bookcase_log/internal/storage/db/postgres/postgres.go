package postgres

import (
	"bookcase_log/models"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	db *sql.DB
}

func New(db *sql.DB) *PostgresStorage {
	return &PostgresStorage{
		db: db,
	}
}

func (s *PostgresStorage) AddLog(lr models.LogRow) error {
	q := `INSERT INTO logs (producer_ts, consumer_ts, topic, message) VALUES($1, $2, $3, $4)`

	err := s.db.QueryRow(
		q,
		lr.Producer_ts,
		lr.Consumer_ts,
		lr.Topic,
		lr.Message,
	)

	if err != nil {
		log.Println("can't add new author: %w", err)
	}

	return nil
}
