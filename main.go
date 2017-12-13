package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"host": c.Request.Host,
	})
}

func uploadHandler(c *gin.Context) {
	uploadedFile, err := c.FormFile("uploadedfile")
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	fileName := uploadedFile.Filename
	log.Printf("File named %v was successfully uploaded.", fileName)
	c.HTML(http.StatusOK, "upload.tmpl", gin.H{
		"fileName": fileName,
	})
}

func main() {
	r := gin.Default()
	r.LoadHTMLFiles(
		"public/index.tmpl",
		"public/upload.tmpl",
	)
	r.GET("/", indexHandler)
	r.POST("/upload", uploadHandler)
	r.Run(":80")
}
