package jwtHelper

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
)

func GetUserId(token string) (interface{}, error) {
	const WHERE = "jwtHelper GetUserId():"
	jwtData, err := GetClaims(token)
	if err != nil {
		return 0, fmt.Errorf("%s невозможно получить данные о пользователе: %s", WHERE, err.Error())
	}

	userIdRaw, ok := jwtData["id"]
	if !ok {
		return 0, fmt.Errorf("%s в jwt отсутствует id пользователя", WHERE)
	}

	return userIdRaw, nil
}

func GetClaims(token string) (jwt.MapClaims, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))

	// парсим токен
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return make(jwt.MapClaims, 0), fmt.Errorf("ошибка парсинга jwt: %s", err.Error())
	}

	res, ok := jwtToken.Claims.(jwt.MapClaims)
	// обязательно используем второе возвращаемое значение ok и проверяем его, потому что
	// если Сlaims вдруг окажется другого типа, мы получим панику
	if !ok {
		return make(jwt.MapClaims, 0), fmt.Errorf("неверный тип jwt.MapClaims")
	}

	return res, nil
}
