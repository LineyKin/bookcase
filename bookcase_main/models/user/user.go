package user

import "bookcase/models/auth"

type User struct {
	Id  int    `json:"id"`
	Jwt string `json:"jwt"`
	auth.AuthData
}

func (u User) GetJWT() (string, error) {
	return u.CalcJWT(u.Id)
}
