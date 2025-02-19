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

func (s *PostgresStorage) GetLogList(limit, offset int) ([]models.LogRow, error) {
	q := `
	SELECT 
		producer_ts,
		message
	FROM logs
	LIMIT $1 OFFSET $2`

	rows, err := s.db.Query(
		q,
		limit,
		offset,
	)

	if err != nil {
		log.Println("PostgresStorage Query error")
		return []models.LogRow{}, err
	}
	defer rows.Close()

	list := []models.LogRow{}
	for rows.Next() {
		lRow := models.LogRow{}
		err := rows.Scan(&lRow.Producer_ts, &lRow.Message)

		if err != nil {
			log.Println("PostgresStorage Scan error")
			return []models.LogRow{}, err
		}

		list = append(list, lRow)
	}

	return list, nil
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
	q := `INSERT INTO logs (user_id, producer_ts, consumer_ts, topic, message) VALUES($1, $2, $3, $4, $5)`

	err := s.db.QueryRow(
		q,
		lr.Id,
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
