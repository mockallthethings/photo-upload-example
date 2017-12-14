package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID   string
	Name string
}

func (h *handler) indexHandler(c *gin.Context) {
	stmt, err := h.db.Prepare("SELECT id,name FROM photos.dim_album")
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	rows, err := stmt.Query()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer rows.Close()

	albums := make([]album, 0)
	for rows.Next() {
		var id string
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		albums = append(albums, album{
			ID:   id,
			Name: name,
		})
	}
	if err = rows.Err(); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// if we got here via a route like /album/:albumID instead of /
	// then show the album contents
	albumID := c.Param("albumID")
	photoIDs := make([]string, 0)
	if albumID != "" {
		photoStmt, err := h.db.Prepare(`
			SELECT
			photo_id
			FROM photos.fact_photo_album
			WHERE album_id = ?;`)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		rows, err := photoStmt.Query(albumID)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var photoID string
			err = rows.Scan(&photoID)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			photoIDs = append(photoIDs, photoID)
		}
		if err = rows.Err(); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	templateParams := gin.H{
		"albums":   albums,
		"albumID":  albumID,
		"photoIDs": photoIDs,
	}

	if albumID != "" {
		templateParams["showUploadForm"] = true
		templateParams["uploadURL"] = "/album/" + albumID + "/upload"
	}

	c.HTML(http.StatusOK, "index.tmpl", templateParams)
}
