package postgres

import (
	"bookcase/models/book"
	"fmt"

	_ "github.com/lib/pq"
)

func getOrderBy(sortedBy, sortType string) string {

	if !(sortType == "asc" || sortType == "desc") {
		sortType = "asc"
	}

	switch sortedBy {
	case "author":
		return fmt.Sprintf("author %s, name", sortType)
	case "name":
		return "name " + sortType
	case "publishingHouse":
		return "publishing_house " + sortType
	case "publishingYear":
		return "year_of_publication " + sortType
	default:
		return "author, name"
	}
}

func (s *PostgresStorage) GetBookList(userId, limit, offset int, sortedBy, sortType string) ([]book.BookUnload, error) {
	q := `
	SELECT
		b.id,
		STRING_AGG(DISTINCT a.last_name || ' ' || a.name, ', ') AS author,
		STRING_AGG(DISTINCT lw.name, '; ') AS name,
		ph.name AS publishing_house,
		b.year_of_publication
	FROM book AS b
	LEFT JOIN publishing_house AS ph ON ph.id = b.publishing_house_id
	LEFT JOIN literary_work AS lw ON lw.id = ANY(b.literary_works)
	LEFT JOIN authors AS a ON a.id = ANY(lw.authors)
	WHERE b.user_id = $3
	GROUP BY b.id, ph.name
	ORDER BY %s
	LIMIT $1 OFFSET $2`

	query := fmt.Sprintf(q, getOrderBy(sortedBy, sortType))

	rows, err := s.db.Query(
		query,
		limit,
		offset,
		userId,
	)

	if err != nil {
		return []book.BookUnload{}, err
	}
	defer rows.Close()

	list := []book.BookUnload{}
	for rows.Next() {
		bRow := book.BookUnload{}
		err := rows.Scan(
			&bRow.Id,
			&bRow.Author,
			&bRow.Name,
			&bRow.PublishingHouse,
			&bRow.PublishingYear,
		)

		if err != nil {
			return []book.BookUnload{}, err
		}

		list = append(list, bRow)
	}

	return list, nil
}

func (s *PostgresStorage) GetBookListTotal(limit, offset int, sortedBy, sortType string) ([]book.BookUnload, error) {
	q := `
	SELECT 
		b.id,
		u.login AS user,
		STRING_AGG(DISTINCT a.last_name || ' ' || a.name, ', ') AS author,
		STRING_AGG(DISTINCT lw.name, '; ') AS name,
		ph.name AS publishing_house,
		b.year_of_publication
	FROM book AS b
	LEFT JOIN publishing_house AS ph ON ph.id = b.publishing_house_id
	LEFT JOIN literary_work AS lw ON lw.id = ANY(b.literary_works)
	LEFT JOIN authors AS a ON a.id = ANY(lw.authors)
	LEFT JOIN users AS u ON u.id = b.user_id
	GROUP BY b.id, ph.name, u.login
	ORDER BY %s
	LIMIT $1 OFFSET $2`

	query := fmt.Sprintf(q, getOrderBy(sortedBy, sortType))

	rows, err := s.db.Query(
		query,
		limit,
		offset,
	)

	if err != nil {
		return []book.BookUnload{}, err
	}
	defer rows.Close()

	list := []book.BookUnload{}
	for rows.Next() {
		bRow := book.BookUnload{}
		err := rows.Scan(
			&bRow.Id,
			&bRow.User,
			&bRow.Author,
			&bRow.Name,
			&bRow.PublishingHouse,
			&bRow.PublishingYear,
		)

		if err != nil {
			return []book.BookUnload{}, err
		}

		list = append(list, bRow)
	}

	return list, nil
}
