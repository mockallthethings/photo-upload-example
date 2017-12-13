package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) uploadHandler(c *gin.Context) {
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
