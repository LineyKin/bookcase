package postgres

import (
	"bookcase/models/book"
	"fmt"

	_ "github.com/lib/pq"
)

func (s *PostgresStorage) AddLiteraryWork(lwName string) (int, error) {
	q := `INSERT INTO literary_work (name) VALUES($1) RETURNING id`
	var id int
	err := s.db.QueryRow(
		q,
		lwName,
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("can't add new literary work: %w", err)
	}

	return id, nil
}

func (s *PostgresStorage) AddPhysicalBook(b *book.BookAdd, userId interface{}) (int, error) {
	q := `INSERT INTO book (user_id, publishing_house_id, year_of_publication) VALUES($1, $2, $3) RETURNING id`
	var id int
	err := s.db.QueryRow(
		q,
		userId,
		b.PublishingHouse.Id,
		b.PublishingYear,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *PostgresStorage) AddPublishingHouse(phName string) (int, error) {
	q := `INSERT INTO publishing_house (name) VALUES($1) RETURNING id`
	var id int
	err := s.db.QueryRow(
		q,
		phName,
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("can't add new publishing house: %w", err)
	}

	return id, nil
}

func (s *PostgresStorage) LinkAuthorAndLiteraryWork(authorId, bookId int) error {
	q := `INSERT INTO author_and_literary_work (author_id, literary_work_id) VALUES($1, $2)`
	_, err := s.db.Exec(
		q,
		authorId,
		bookId,
	)

	if err != nil {
		return fmt.Errorf("can't link author and literary work: %w", err)
	}

	return nil
}

func (s *PostgresStorage) LinkBookAndLiteraryWork(lwId, bookId int) error {
	q := `INSERT INTO book_and_literary_work (literary_work_id, book_id) VALUES($1, $2)`
	_, err := s.db.Exec(
		q,
		lwId,
		bookId,
	)

	if err != nil {
		return fmt.Errorf("can't link literary work and book: %w", err)
	}

	return nil
}
