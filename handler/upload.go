package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mockallthethings/photo-upload-example/s3"
)

func (h *handler) uploadHandler(c *gin.Context) {
	albumID := c.Param("albumID")
	uploadedFile, err := c.FormFile("uploadedfile")
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	photoID, err := uuid.NewRandom()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	f, err := uploadedFile.Open()
	fileName := uploadedFile.Filename
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	defaultBucket := "mockallthethings-example"
	idStr := photoID.String()
	ui := &s3.UploadInput{
		Bucket: &defaultBucket,
		Key:    &idStr,
		Body:   f,
	}
	_, err = h.S3.Upload(ui)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	stmt, err := h.db.Prepare(`
INSERT INTO photos.fact_photo_album (
	album_id,
	photo_id
) VALUES (?, ?);`)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	_, err = stmt.Exec(
		albumID,
		photoID,
	)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	log.Printf("File named %v was successfully uploaded.", fileName)

	c.Redirect(http.StatusFound, "/album/"+albumID+"/view")
}
