package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"host": c.Request.Host,
	})
}

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("public/index.tmpl")
	r.GET("/", indexHandler)
	r.Run(":80")
}
