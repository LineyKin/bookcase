package handlers

import (
	"bookcase/models"
	"bookcase/models/author"
	"bookcase/models/book"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

func GetObjectByPath(c *gin.Context) models.UserLogInterface {
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

func (ctrl *Controller) getStringFromContext(key string, c *gin.Context) (string, error) {
	strRaw, ok := c.Get(key)
	if !ok {
		log.Println("no key: ", key)
		return "", fmt.Errorf("no key: %s", key)
	}

	log.Println(strRaw, reflect.TypeOf(strRaw))

	str, ok := strRaw.(string)
	if !ok {
		return "", fmt.Errorf("%s не является string", key)
	}

	return str, nil
}

func (ctrl *Controller) getUserLogFromContext(c *gin.Context) (models.UserLog, error) {
	ulRaw, ok := c.Get(USER_LOG_KEY)
	if !ok {
		log.Println("no key: ", USER_LOG_KEY)
		return models.UserLog{}, fmt.Errorf("no key: %s", USER_LOG_KEY)
	}

	ul, ok := ulRaw.(models.UserLog)
	if !ok {
		return models.UserLog{}, fmt.Errorf("%s не является типом models.UserLog", USER_LOG_KEY)
	}

	return ul, nil
}

func (ctrl *Controller) getUserLogin(c *gin.Context) (string, error) {
	return ctrl.getStringFromContext(LOGIN_KEY, c)
}
