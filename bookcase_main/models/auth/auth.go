package auth

import (
	"bookcase/models"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
)

const REGISTER_ACTION = 1
const LOGIN_ACTION = 2

type AuthData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Action   int    `json:"action,omitempty"`
}

func (a AuthData) NewLog() models.UserLog {
	ul := models.NewUserLog()
	switch a.Action {
	case LOGIN_ACTION:
		ul.Message = fmt.Sprintf("Пользователь %s авторизовался на сайте", a.Login)
	case REGISTER_ACTION:
		ul.Message = fmt.Sprintf("Пользователь %s зарегистрировался на сайте", a.Login)
	default:
		ul.Message = fmt.Sprintf("Неизвестное действие пользователя %s", a.Login)
	}

	return ul
}

func (a AuthData) CalcPwdHash(pwd string) string {
	saltedPwd := fmt.Sprintf("%s%s%s", os.Getenv("SALT_1"), pwd, os.Getenv("SALT_2"))
	result := sha256.Sum256([]byte(saltedPwd))

	return hex.EncodeToString(result[:])
}

func (a *AuthData) HashPwd() {
	a.Password = a.CalcPwdHash(a.Password)
}

func (a *AuthData) CalcJWT(userId int) (string, error) {
	claims := jwt.MapClaims{
		"id":    userId,
		"login": a.Login,
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
