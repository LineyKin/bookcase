package sqlite

import (
	"bookcase/models/auth"
	"bookcase/models/author"
	"bookcase/models/book"
	u "bookcase/models/user"
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type SqliteStorage struct {
	db *sql.DB
}

func New(db *sql.DB) *SqliteStorage {
	return &SqliteStorage{
		db: db,
	}
}

func (s *SqliteStorage) GetUserByAuthLogin(a auth.AuthData) (u.User, error) {
	q := `SELECT * FROM users 
			WHERE login=$1`

	var user u.User

	err := s.db.QueryRow(q,
		a.Login,
		a.Password,
	).Scan(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *SqliteStorage) GetBookCount() (int, error) {
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

func (s *SqliteStorage) GetPublishingHouseList() ([]book.PublishingHouse, error) {
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

func (s *SqliteStorage) GetAuthorByName(a author.Author) ([]int, error) {
	q := `SELECT id FROM authors 
			WHERE name=:name
			AND father_name=:father_name
			AND last_name=:last_name`

	rows, err := s.db.Query(q,
		sql.Named("name", a.Name),
		sql.Named("father_name", a.FatherName),
		sql.Named("last_name", a.LastName),
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

		list = append(list, int(id))
	}

	return list, nil
}

func (s *SqliteStorage) GetAuthorList() ([]author.Author, error) {
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

func (s *SqliteStorage) AddAuthor(a author.Author) (int, error) {
	q := `INSERT INTO authors (name, father_name, last_name) VALUES(:name, :father_name, :last_name)`

	res, err := s.db.Exec(
		q,
		sql.Named("name", a.Name),
		sql.Named("father_name", a.FatherName),
		sql.Named("last_name", a.LastName),
	)

	if err != nil {
		return 0, fmt.Errorf("can't add new author: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
