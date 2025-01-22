package postgres

import (
	"bookcase_log/models"
	"database/sql"
	"log"
	"time"

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

func (s *PostgresStorage) GetLogCount() (int, error) {
	q := `SELECT COUNT(*) FROM logs`
	row, err := s.db.Query(q)
	if err != nil {
		return 0, err
	}
	defer row.Close()

	var count int

	row.Next()
	err = row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (s *PostgresStorage) GetLatestLogTimestamp() time.Time {
	var ts time.Time
	q := "SELECT MAX(producer_ts) FROM logs"
	err := s.db.QueryRow(q).Scan(&ts)

	if err != nil {
		log.Println("can't get latest timestamp: %w", err)
	}

	return ts
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
