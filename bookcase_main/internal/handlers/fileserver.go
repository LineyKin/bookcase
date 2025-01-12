package handlers

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

const webDir = "web"

func (ctrl *Controller) FileServer(c *gin.Context) {
	log.Println("we are in fileserver number 44222")
	filePath := filepath.Join(webDir, strings.TrimPrefix(c.Request.URL.Path, "/"))
	c.File(filePath)
}
