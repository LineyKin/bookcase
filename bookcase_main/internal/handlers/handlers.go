package handlers

import (
	"bookcase/internal/kafka"
	"bookcase/internal/service"

	"github.com/gin-gonic/gin"
)

type HandlersInterface interface {
	AddAuthor(c *gin.Context)
	GetAuthorList(c *gin.Context)
	AddBook(c *gin.Context)
	GetPublishingHouseList(c *gin.Context)
	GetBookCount(c *gin.Context)
	GetBookList(c *gin.Context)
	FileServer(c *gin.Context)
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type Handlers struct {
	HandlersInterface
}

func New(services *service.Service, kp *kafka.Producer) *Handlers {
	return &Handlers{
		HandlersInterface: NewController(services, kp),
	}
}
