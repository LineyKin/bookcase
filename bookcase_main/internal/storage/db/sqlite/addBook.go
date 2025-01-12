package sqlite

import (
	"bookcase/models/book"
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func (s *SqliteStorage) AddLiteraryWork(lwName string) (int, error) {
	q := `INSERT INTO literary_work (name) VALUES(:name)`
	res, err := s.db.Exec(
		q,
		sql.Named("name", lwName),
	)

	if err != nil {
		return 0, fmt.Errorf("can't add new literary work: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *SqliteStorage) AddPhysicalBook(b *book.BookAdd) (int, error) {
	q := `INSERT INTO book (publishing_house_id, year_of_publication) VALUES(:ph_id, :yop)`
	res, err := s.db.Exec(
		q,
		sql.Named("ph_id", b.PublishingHouse.Id),
		sql.Named("yop", b.PublishingYear),
	)

	if err != nil {
		return 0, err
	}

	bookId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(bookId), nil
}

func (s *SqliteStorage) AddPublishingHouse(phName string) (int, error) {
	q := `INSERT INTO publishing_house (name) VALUES(:name)`
	res, err := s.db.Exec(
		q,
		sql.Named("name", phName),
	)

	if err != nil {
		return 0, fmt.Errorf("can't add new publishing house: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *SqliteStorage) LinkAuthorAndLiteraryWork(authorId, bookId int) error {
	q := `INSERT INTO author_and_literary_work (author_id, literary_work_id) VALUES(:a_id, :b_id)`
	_, err := s.db.Exec(
		q,
		sql.Named("a_id", authorId),
		sql.Named("b_id", bookId),
	)

	if err != nil {
		return fmt.Errorf("can't link author and literary work: %w", err)
	}

	return nil
}

func (s *SqliteStorage) LinkBookAndLiteraryWork(lwId, bookId int) error {
	q := `INSERT INTO book_and_literary_work (literary_work_id, book_id) VALUES(:lw_id, :b_id)`
	_, err := s.db.Exec(
		q,
		sql.Named("lw_id", lwId),
		sql.Named("b_id", bookId),
	)

	if err != nil {
		return fmt.Errorf("can't link literary work and book: %w", err)
	}

	return nil
}
