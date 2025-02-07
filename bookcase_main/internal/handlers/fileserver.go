package handlers

import (
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
	filePath := filepath.Join(WEB_DIR, strings.TrimPrefix(urlPath, "/"))
	c.File(filePath)
}
