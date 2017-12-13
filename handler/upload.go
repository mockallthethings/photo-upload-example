package handler

import (
	"fmt"
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
	albumName := c.PostForm("albumname")
	if albumName == "" {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("Album name is required"))
		return
	}

	stmt, err := h.db.Prepare(`
INSERT INTO photos.albums (
	album_name,
	photo_name
) VALUES (?, ?);`)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	fileName := uploadedFile.Filename
	_, err = stmt.Exec(
		albumName,
		fileName,
	)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	log.Printf("File named %v was successfully uploaded.", fileName)

	c.HTML(http.StatusOK, "upload.tmpl", gin.H{
		"fileName": fileName,
	})
}
