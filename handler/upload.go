package handler

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mockallthethings/photo-upload-example/s3"
)

func (h *handler) uploadHandler(c *gin.Context) {
	albumName := c.PostForm("albumname")
	if albumName == "" {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("Album name is required"))
		return
	}

	uploadedFile, err := c.FormFile("uploadedfile")
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	f, err := uploadedFile.Open()
	fileName := uploadedFile.Filename
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	defaultBucket := "mockallthethings-example"
	ui := &s3.UploadInput{
		Bucket: &defaultBucket,
		Key:    &fileName,
		Body:   f,
	}
	_, err = h.S3.Upload(ui)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	stmt, err := h.db.Prepare(`
INSERT INTO photos.albums (
	album_name,
	photo_id
) VALUES (?, ?);`)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	photoID, err := uuid.NewRandom()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	_, err = stmt.Exec(
		albumName,
		photoID,
	)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	log.Printf("File named %v was successfully uploaded.", fileName)

	c.HTML(http.StatusOK, "upload.tmpl", gin.H{
		"fileName": fileName,
		"albumURL": "/view?album=" + url.QueryEscape(albumName),
	})
}
