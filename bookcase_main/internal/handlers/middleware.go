package handlers

import (
	"bookcase/internal/lib/jwtHelper"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const COOKIE_JWT_KEY = "bookcase_jwt"
const USER_ID_KEY = "user_id"

// Middleware для проверки JWT
func (ctrl *Controller) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie(COOKIE_JWT_KEY)
		if err != nil {
			c.Abort()

			// если пользователь не зарегистрирован/залогинен, отправляем его на страницу auth
			c.Redirect(http.StatusFound, "/auth")
			return
		}

		userId, err := jwtHelper.GetUserId(tokenString)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("ошибка получения id пользователя: %s", err)})
			c.Abort()
			return
		}

		log.Println("userId we set", userId)
		c.Set(USER_ID_KEY, userId)
		c.Next()
	}
}
