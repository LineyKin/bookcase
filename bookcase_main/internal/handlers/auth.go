package handlers

import (
	"bookcase/models/auth"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const COOKIE_JWT_AGE = 604800 // 7 дней в секундах
const HTTP_ONLY = true
const SECURED = false
const HOST = "localhost"

func (ctrl *Controller) Register(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Метод не поддерживается"})
		return
	}

	var authData auth.AuthData
	if err := c.BindJSON(&authData); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка десериализации JSON"})
		return
	}

	// 1. проверяем, есть ли вообще такой пользователь в базе
	_, found, err := ctrl.service.Identify(authData)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка идентификации пользователя"})
		return
	}

	if found {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь с таким логином уже существует"})
		return
	}

	// сохраняем пользователя в систему и получаем jwt
	jwt, err := ctrl.service.AddNewUser(authData)
	if err != nil {
		log.Println("Ошибка добавления нового пользователя", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка добавления нового пользователя"})
		return
	}

	setJWTCookie(c, jwt)
	c.JSON(http.StatusCreated, gin.H{})
}

func (ctrl *Controller) Login(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Метод не поддерживается"})
		return
	}

	var authData auth.AuthData
	if err := c.BindJSON(&authData); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка десериализации JSON"})
		return
	}

	// 1. проверяем, есть ли вообще такой пользователь в базе
	user, found, err := ctrl.service.Identify(authData)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка идентификации пользователя"})
		return
	}

	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден. Зарегистрируйтесь"})
		return
	}

	// 2. тут сверяем пароль
	if user.Password != authData.CalcPwdHash(authData.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный пароль"})
		return
	}

	// 3. возвращаем jwt
	jwt, err := user.GetJWT()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания jwt"})
		return
	}

	setJWTCookie(c, jwt)
	c.JSON(http.StatusAccepted, gin.H{})
}

func setJWTCookie(c *gin.Context, jwt string) {
	c.SetCookie(COOKIE_JWT_KEY, jwt, COOKIE_JWT_AGE, "/", HOST, SECURED, HTTP_ONLY)
}
