package handlers

import (
	"bookcase/internal/lib/jwtHelper"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const COOKIE_JWT_KEY = "bookcase_jwt"
const USER_ID_KEY = "user_id"
const LOGIN_KEY = "login"
const USER_LOG_KEY = "user_log"

type Middleware interface {
	AuthMW() gin.HandlerFunc
	LogMW() gin.HandlerFunc
}

func (ctrl *Controller) LogMW() gin.HandlerFunc {
	return func(c *gin.Context) {
		userLog, err := ctrl.getUserLogFromContext(c)
		if err != nil {
			log.Println("ошибка получения лога из контекста в LogMW()", err)
			c.Abort()
			return
		}

		userId, err := getUserId(c)
		if err != nil {
			log.Println("ошибка получения id пользователя из контекста в LogMW()", err)
			c.Abort()
			return
		}

		login, err := ctrl.getUserLogin(c)
		if err != nil {
			log.Println("ошибка получения логина из контекста в LogMW()", err)
			c.Abort()
			return
		}

		userLog.Id = userId
		userLog.Login = login

		// если продюсер кафки не активен, завершаем хэндлер
		if ctrl.kp == nil {
			log.Println("лог не передан в кафку из-за её неактивности")
			c.Abort()
			return
		}

		// data for kafka
		logInBytes, err := json.Marshal(userLog)
		if err != nil {
			log.Println("can't prepare data for kafka: ", err)
			c.Abort()
			return
		}

		// Send the bytes to kafka
		err = ctrl.kp.PushLogToQueue("bookcase_log", logInBytes)
		if err != nil {
			log.Println("can't send data to kafka: ", err)
			c.Abort()
			return
		}

		c.Abort()
	}
}

// Middleware для проверки JWT
func (ctrl *Controller) AuthMW() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie(COOKIE_JWT_KEY)
		if err != nil {
			c.Abort()

			// если пользователь не зарегистрирован/залогинен, отправляем его на страницу auth
			c.Redirect(http.StatusFound, "/auth")
			return
		}

		userInfo, err := jwtHelper.GetUserInfo(tokenString)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("ошибка получения данных пользователя из jwt: %s", err)})
			c.Abort()
			return
		}

		c.Set(USER_ID_KEY, userInfo[jwtHelper.USER_ID_INDEX])
		c.Set(LOGIN_KEY, userInfo[jwtHelper.USER_LOGIN_INDEX])
		c.Next()
	}
}
