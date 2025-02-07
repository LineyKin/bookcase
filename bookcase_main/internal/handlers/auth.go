package handlers

import (
	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) Register(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Пользователь заригистрирован"})
}
