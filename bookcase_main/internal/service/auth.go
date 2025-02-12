package service

import (
	"bookcase/models/auth"
	u "bookcase/models/user"
	"log"
)

type AuthInterface interface {
	Identify(authData auth.AuthData) (u.User, bool, error)
	AddNewUser(authData auth.AuthData) (userJWT string, err error)
}

func (s *bookcaseService) AddNewUser(authData auth.AuthData) (userJWT string, err error) {
	authData.HashPwd()
	id, err := s.storage.AddNewUser(authData)
	if err != nil {
		log.Println("Сервис: ошибка добавления нового пользователя:", err)
		return "", err
	}

	return authData.CalcJWT(id)
}

func (s *bookcaseService) Identify(data auth.AuthData) (u.User, bool, error) {
	user, err := s.storage.GetUserByAuthLogin(data)

	if err != nil {
		log.Println("Сервис: ошибка идентификации пользователя:", err)
		return user, false, err
	}

	if user.Id == 0 {
		return user, false, nil
	}

	return user, true, nil
}
