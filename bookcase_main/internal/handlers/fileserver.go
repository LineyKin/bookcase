package handlers

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

const WEB_DIR = "web"
const EXTENSION = "html"

func (ctrl *Controller) FileServer(c *gin.Context) {
	urlPath := c.Request.URL.Path
	if urlPath != "/" {
		urlPath += "." + EXTENSION
	}
	uId, ok := c.Get("user_id")
	if !ok {
		log.Println("no user_id key")
	}
	log.Println("fileserver user id", uId)
	filePath := filepath.Join(WEB_DIR, strings.TrimPrefix(urlPath, "/"))
	c.File(filePath)
}
