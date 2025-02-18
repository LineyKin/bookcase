package postgres

import (
	"bookcase/models/auth"
	u "bookcase/models/user"
	"database/sql"
	"fmt"
	"log"
)

const DUPLICATE_ERROR = `pq: duplicate key value violates unique constraint "idx_login_unique"`

func (s *PostgresStorage) AddNewUser(data auth.AuthData) (int, error) {
	q := `INSERT INTO users (login, password) VALUES($1, $2) RETURNING id`
	var id int
	err := s.db.QueryRow(
		q,
		data.Login,
		data.Password,
	).Scan(&id)

	if err != nil {
		log.Printf("can't add new user: %s", err)

		if err.Error() == DUPLICATE_ERROR {
			return 0, fmt.Errorf("пользователь с логином %s уже существует", data.Login)
		}

		return id, err
	}

	return id, nil
}

func (s *PostgresStorage) GetUserByAuthLogin(a auth.AuthData) (u.User, error) {
	q := `SELECT * FROM users WHERE login=$1`

	var user u.User

	err := s.db.QueryRow(q, a.Login).Scan(&user.Id, &user.Login, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, nil
		}

		return user, err
	}

	return user, nil
}
