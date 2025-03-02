package postgres

import (
	"bookcase/models/book"
	"fmt"

	_ "github.com/lib/pq"
)

func (s *PostgresStorage) AddBookWithNewPublishingHouse(b *book.BookAdd, userId interface{}) error {

	// Начинаем транзакцию
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	// формируем запрос
	q := `
	WITH 
		lwId AS (
	    	INSERT INTO literary_work (name, authors)
	    	VALUES %s
	    	RETURNING id
		),
		phId AS (
	    	INSERT INTO publishing_house (name)
	    	VALUES ($2)
	    	RETURNING id
		)
	INSERT INTO book (
	    user_id, 
	    publishing_house_id, 
	    year_of_publication, 
	    literary_works
	)
	VALUES(
		$1,
		(SELECT id FROM phId),
		$3,
		(SELECT array_agg(id) FROM lwId)
	);`

	q = fmt.Sprintf(q, b.GetLWInsertion())

	_, err = tx.Exec(
		q,
		userId,
		b.PublishingHouse.Name,
		b.PublishingYear,
		b.GetAuthorIdsAsArrayForPG(), // $4 authors, массив id (прилетит в %s)
	)

	if err != nil {
		// Откатываем транзакцию в случае ошибки
		tx.Rollback()
		return fmt.Errorf("can't add new book with new publishing house: %w", err)
	}

	// Фиксируем транзакцию
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (s *PostgresStorage) AddBook(b *book.BookAdd, userId interface{}) error {
	// Начинаем транзакцию
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	// формируем запрос
	q := `
	WITH lwId AS (
	    INSERT INTO literary_work (name, authors)
	    VALUES %s
	    RETURNING id
	)
	INSERT INTO book (
	    user_id, 
	    publishing_house_id, 
	    year_of_publication, 
	    literary_works
	)
	SELECT 
	    $1, 
	    $2, 
	    $3, 
	    array_agg(id)
	FROM lwId;`

	q = fmt.Sprintf(q, b.GetLWInsertion())

	_, err = tx.Exec(
		q,
		userId,
		b.PublishingHouse.Id,
		b.PublishingYear,
		b.GetAuthorIdsAsArrayForPG(), // $4 authors, массив id (прилетит в %s)
	)

	if err != nil {
		// Откатываем транзакцию в случае ошибки
		tx.Rollback()
		return fmt.Errorf("can't add new book: %w", err)
	}

	// Фиксируем транзакцию
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
