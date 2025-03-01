package handlers

import (
	"bookcase/models"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) getStringFromContext(key string, c *gin.Context) (string, error) {
	strRaw, ok := c.Get(key)
	if !ok {
		log.Println("no key: ", key)
		return "", fmt.Errorf("no key: %s", key)
	}

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

func (ctrl *Controller) getUserId(c *gin.Context) (int, error) {
	userIdRaw, ok := c.Get(USER_ID_KEY)
	if !ok {
		log.Println("no key: ", USER_ID_KEY)
		return 0, fmt.Errorf("no key: %s", USER_ID_KEY)
	}

	userId, ok := userIdRaw.(float64)
	if !ok {
		return 0, fmt.Errorf("userId не является float64")
	}

	return int(userId), nil
}
