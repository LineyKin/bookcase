package handlers

import (
	"github.com/gin-gonic/gin"
)

type authData struct {
	login    string
	password string
}

func (ctrl *Controller) Register(c *gin.Context) {
	// 1. Проверка на существование
	// 1. Если пользователь уже существует, возврат в исходное с ошибкой 400
	c.JSON(201, gin.H{"message": "Пользователь зарегистрирован"})
}

func (ctrl *Controller) Login(c *gin.Context) {
	// 1 идентификация
	// 1.1 если не пройдена (такого логина нет в БД) - перевод на регистрацию с ошибкой 404

	// 2. аутентификация, возвращение jwt клиенту
	c.JSON(200, gin.H{"message": "Пользователь вошёл"})
}
