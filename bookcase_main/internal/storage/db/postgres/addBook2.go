package postgres

import (
	"bookcase/models/book"
	"fmt"
)

func (s *PostgresStorage) AddBookWithNewPublishingHouse(b *book.BookAdd, userId interface{}) error {
	return nil
}

func (s *PostgresStorage) AddBook(b *book.BookAdd, userId interface{}) error {
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

	_, err := s.db.Exec(
		q,
		userId,
		b.PublishingHouse.Id,
		b.PublishingYear,
		b.GetAuthorIdsAsArrayForPG(), // $4 authors, массив id (прилетит в %s)
	)

	if err != nil {
		return fmt.Errorf("can't add new book: %w", err)
	}

	return nil
}
