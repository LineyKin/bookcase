package jwtHelper

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
)

const USER_ID_INDEX = 0
const USER_LOGIN_INDEX = 1

func GetUserInfo(token string) ([2]any, error) {
	const WHERE = "jwtHelper GetUserId():"
	jwtData, err := GetClaims(token)
	userInfo := [2]any{}
	if err != nil {
		return userInfo, fmt.Errorf("%s невозможно получить данные о пользователе: %s", WHERE, err.Error())
	}

	userIdRaw, ok := jwtData["id"]
	if !ok {
		return userInfo, fmt.Errorf("%s в jwt отсутствует id пользователя", WHERE)
	}

	userInfo[USER_ID_INDEX] = userIdRaw

	loginRaw, ok := jwtData["login"]
	if !ok {
		return userInfo, fmt.Errorf("%s в jwt отсутствует login пользователя", WHERE)
	}

	userInfo[USER_LOGIN_INDEX] = loginRaw

	return userInfo, nil
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
