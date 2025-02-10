package user

import "bookcase/models/auth"

type User struct {
	Id  int    `json:"id"`
	Jwt string `json:"jwt"`
	auth.AuthData
}
