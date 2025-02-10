package postgres

import (
	"bookcase/models/auth"
	u "bookcase/models/user"
	"database/sql"
	"fmt"
)

func (s *PostgresStorage) AddNewUser(data auth.AuthData) (int, error) {
	q := `INSERT INTO users (login, password) VALUES($1, $2) RETURNING id`
	var id int
	err := s.db.QueryRow(
		q,
		data.Login,
		data.Password,
	).Scan(&id)

	if err != nil {
		return id, fmt.Errorf("can't add new user: %w", err)
	}

	return id, nil
}

func (s *PostgresStorage) GetUserByAuthLogin(a auth.AuthData) (u.User, error) {
	q := `SELECT * FROM users WHERE login=$1`

	var user u.User

	err := s.db.QueryRow(q, a.Login).Scan(&user)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, nil
		}

		return user, err
	}

	return user, nil
}
