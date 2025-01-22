package handlers

import (
	"bookcase_log/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service service.ServiceInterface
}

func NewController(service service.ServiceInterface) *Controller {
	return &Controller{
		service: service,
	}
}

func (ctrl *Controller) GetLogList(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Метод не поддерживается"})
		return
	}

	limitString := c.Query("limit")
	limit, err := strconv.Atoi(limitString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	offsetString := c.Query("offset")
	offset, err := strconv.Atoi(offsetString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logList, err := ctrl.service.GetLogList(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"log_list": logList})
}

func (ctrl *Controller) GetLogCount(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Метод не поддерживается"})
		return
	}

	count, err := ctrl.service.GetLogCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"count": count})
}
