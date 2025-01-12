package sqlite

import (
	"bookcase/models/book"
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func getOrderBy(sortedBy, sortType string) string {

	if !(sortType == "asc" || sortType == "desc") {
		sortType = "asc"
	}

	switch sortedBy {
	case "author":
		return fmt.Sprintf("a.last_name %s, lw.name", sortType)
	case "name":
		return "lw.name " + sortType
	case "publishingHouse":
		return "ph.name " + sortType
	case "publishingYear":
		return "b.year_of_publication " + sortType
	default:
		return "a.last_name, lw.name"
	}
}

func (s *SqliteStorage) GetBookList(limit, offset int, sortedBy, sortType string) ([]book.BookUnload, error) {
	q := `
	SELECT
 		b.id AS id,
 		GROUP_CONCAT(DISTINCT IFNULL(a.last_name || ' ' || a.name, '-')) AS author,
 		GROUP_CONCAT(DISTINCT lw.name) AS name,
 		ph.name AS publishing_house,
 		b.year_of_publication
	FROM book AS b
	LEFT JOIN publishing_house AS ph ON ph.id = b.publishing_house_id
	LEFT JOIN book_and_literary_work AS blw ON blw.book_id = b.id
	LEFT JOIN literary_work AS lw ON lw.id = blw.literary_work_id
	LEFT JOIN author_and_literary_work AS alw ON alw.literary_work_id = lw.id
	LEFT JOIN authors AS a ON a.id = alw.author_id
	GROUP BY b.id
	ORDER BY %s
	LIMIT :limit OFFSET :offset;`

	query := fmt.Sprintf(q, getOrderBy(sortedBy, sortType))

	rows, err := s.db.Query(
		query,
		sql.Named("limit", limit),
		sql.Named("offset", offset),
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
