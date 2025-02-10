package service

import (
	"bookcase/models/auth"
	u "bookcase/models/user"
	"fmt"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
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

	claims := jwt.MapClaims{
		"id":    id,
		"login": authData.Login,
	}

	// создаём jwt и указываем payload
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := []byte(os.Getenv("JWT_SECRET"))

	// получаем подписанный токен
	signedToken, err := jwtToken.SignedString(secret)
	if err != nil {
		return "", fmt.Errorf("failed to sign jwt: %s", err)
	}

	return signedToken, nil
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
