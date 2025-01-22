package handlers

import (
	"bookcase_log/internal/service"

	"github.com/gin-gonic/gin"
)

type HandlersInterface interface {
	GetLogCount(c *gin.Context)
	FileServer(c *gin.Context)
	GetLogList(c *gin.Context)
}

type Handlers struct {
	HandlersInterface
}

func New(services *service.Service) *Handlers {
	return &Handlers{
		HandlersInterface: NewController(services),
	}
}
