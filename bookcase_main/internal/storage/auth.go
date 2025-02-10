package storage

import (
	"bookcase/models/auth"
	u "bookcase/models/user"
)

type AuthInterface interface {
	GetUserByAuthLogin(data auth.AuthData) (u.User, error)
	AddNewUser(data auth.AuthData) (int, error)
}
