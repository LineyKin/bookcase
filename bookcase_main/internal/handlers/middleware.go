package handlers

import (
	"bookcase/internal/lib/jwtHelper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const COOKIE_JWT = "bookcase_jwt"
const USER_ID_KEY = "user_id"

// Middleware для проверки JWT
func (ctrl *Controller) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie(COOKIE_JWT)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не зарегистрирован"})
			c.Abort()
			return
		}

		userId, err := jwtHelper.GetUserId(tokenString)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("ошибка получения id пользователя: %s", err)})
			c.Abort()
			return
		}
		c.Set(USER_ID_KEY, userId)
		c.Next()
	}
}
