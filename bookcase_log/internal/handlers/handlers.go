package handlers

import "github.com/gin-gonic/gin"

type HandlersInterface interface {
	FileServer(c *gin.Context)
}

type Handlers struct {
	HandlersInterface
}

func New() *Handlers {
	return &Handlers{}
}
