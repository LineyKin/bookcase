package user

import "bookcase/models/auth"

type User struct {
	Id int `json:"id"`
	auth.AuthData
}
