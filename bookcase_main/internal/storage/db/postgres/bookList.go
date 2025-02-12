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
		STRING_AGG(DISTINCT a.last_name || ' ' || a.name, ',') AS author,
		STRING_AGG(DISTINCT lw.name, ',') AS name,
		ph.name AS publishing_house,
		b.year_of_publication
	FROM book AS b
	LEFT JOIN publishing_house AS ph ON ph.id = b.publishing_house_id
	LEFT JOIN book_and_literary_work AS blw ON blw.book_id = b.id
	LEFT JOIN literary_work AS lw ON lw.id = blw.literary_work_id
	LEFT JOIN author_and_literary_work AS alw ON alw.literary_work_id = lw.id
	LEFT JOIN authors AS a ON a.id = alw.author_id
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
		err := rows.Scan(&bRow.Id, &bRow.Author, &bRow.Name, &bRow.PublishingHouse, &bRow.PublishingYear)

		if err != nil {
			return []book.BookUnload{}, err
		}

		list = append(list, bRow)
	}

	return list, nil
}
