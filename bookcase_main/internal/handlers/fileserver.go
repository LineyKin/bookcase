package handlers

import (
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

const webDir = "web"

func (ctrl *Controller) FileServer(c *gin.Context) {
	filePath := filepath.Join(webDir, strings.TrimPrefix(c.Request.URL.Path, "/"))
	c.File(filePath)
}
