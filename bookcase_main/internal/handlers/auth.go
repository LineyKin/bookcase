package handlers

import (
	"bookcase/models/auth"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	user, found, err := ctrl.service.Identify(authData)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка идентификации пользователя"})
		return
	}

	if found {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь с таким логином уже существует"})
		return
	}

	_ = user

	c.JSON(201, gin.H{"message": "Пользователь зарегистрирован"})
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
	if user.Password != "111" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный пароль"})
		return

	}

	// 3. возвращаем jwt

	c.JSON(http.StatusOK, gin.H{"message": "Пользователь вошёл"})
}
