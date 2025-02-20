package handlers

import (
	"bookcase/internal/lib/jwtHelper"
	"bookcase/models"
	"bookcase/models/author"
	"bookcase/models/book"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const COOKIE_JWT_KEY = "bookcase_jwt"
const USER_ID_KEY = "user_id"

type Middleware interface {
	AuthMW() gin.HandlerFunc
	LogMW() gin.HandlerFunc
}

func getObjectByPath(c *gin.Context) models.UserLogInterface {
	var err error

	defer func() {
		if err != nil {
			errOutput := fmt.Errorf("ошибка десериализации JSON во время логирования: %s", err)
			log.Println(errOutput)
			c.JSON(http.StatusBadRequest, gin.H{"error": errOutput})
			c.Abort()
			return
		}
	}()

	switch c.Request.URL.Path {
	case ADD_AUTHOR_URL:
		var author author.Author
		err = c.BindJSON(&author)
		return author
	case ADD_BOOK_URL:
		var book book.BookAdd
		err = c.BindJSON(&book)
		return book
	default:
		return nil
	}
}

func (ctrl *Controller) LogMW() gin.HandlerFunc {
	return func(c *gin.Context) {

		// копируем тело запроса
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		// Возвращаем в контекст.
		// Оно у нас опустошилось после копирования.
		c.Request.Body = io.NopCloser(bytes.NewReader(body))

		userLog := getObjectByPath(c).NewLog()
		userId, _ := getUserId(c)
		userLog.Id = userId

		// если продюсер кафки не активен, завершаем хэндлер
		if ctrl.kp == nil {
			log.Println("лог не передан в кафку из-за её неактивности")
			c.Next()
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

		// Снова возвращаем тело в контекст.
		// Оно у нас опустошилось после чтения в getObjectByPath()
		c.Request.Body = io.NopCloser(bytes.NewReader(body))
		c.Next()
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
