package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"host": c.Request.Host,
	})
}
