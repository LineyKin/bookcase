package postgres

import (
	"bookcase/models/author"
	"bookcase/models/book"
	"database/sql"
	"fmt"
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

func (s *PostgresStorage) GetBookCount() (int, error) {
	q := `SELECT COUNT(*) FROM book`
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

func (s *PostgresStorage) GetPublishingHouseList() ([]book.PublishingHouse, error) {
	q := `SELECT * FROM publishing_house ORDER BY name`

	rows, err := s.db.Query(q)
	if err != nil {
		return []book.PublishingHouse{}, err
	}
	defer rows.Close()

	list := []book.PublishingHouse{}
	for rows.Next() {
		ph := book.PublishingHouse{}
		err := rows.Scan(&ph.Id, &ph.Name)

		if err != nil {
			return []book.PublishingHouse{}, err
		}

		list = append(list, ph)
	}

	return list, nil
}

func (s *PostgresStorage) GetAuthorByName(a author.Author) ([]int, error) {
	q := `SELECT id FROM authors 
			WHERE name=$1
			AND father_name=$2
			AND last_name=$3`

	rows, err := s.db.Query(q,
		a.Name,
		a.FatherName,
		a.LastName,
	)
	list := []int{}
	if err != nil {
		return list, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			return list, err
		}

		list = append(list, id)
	}

	return list, nil
}

func (s *PostgresStorage) GetAuthorList() ([]author.Author, error) {
	q := `SELECT * FROM authors ORDER BY last_name`

	rows, err := s.db.Query(q)
	if err != nil {
		return []author.Author{}, err
	}
	defer rows.Close()

	list := []author.Author{}
	for rows.Next() {
		authorRow := author.Author{}
		err := rows.Scan(&authorRow.Id, &authorRow.Name, &authorRow.FatherName, &authorRow.LastName)

		if err != nil {
			return []author.Author{}, err
		}

		list = append(list, authorRow)
	}

	return list, nil
}

func (s *PostgresStorage) AddAuthor(a author.Author) (int, error) {
	q := `INSERT INTO authors (name, father_name, last_name) VALUES($1, $2, $3) RETURNING id`

	var id int
	err := s.db.QueryRow(
		q,
		a.Name,
		a.FatherName,
		a.LastName,
	).Scan(&id)

	if err != nil {
		log.Println("can't add new author: %w", err)
		return 0, fmt.Errorf("can't add new author: %w", err)
	}

	return id, nil
}
