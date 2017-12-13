package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) viewHandler(c *gin.Context) {
	albumName := c.Query("album")
	if albumName == "" {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("Album name is required"))
		return
	}

	stmt, err := h.db.Prepare(`
SELECT
  photo_name
FROM photos.albums
WHERE album_name = ?;`)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	rows, err := stmt.Query(albumName)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer rows.Close()

	photoNames := make([]string, 0)
	for rows.Next() {
		var photoName string
		err := rows.Scan(&photoName)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		photoNames = append(photoNames, photoName)
	}
	if err = rows.Err(); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.HTML(http.StatusOK, "view.tmpl", gin.H{
		"albumName":  albumName,
		"photoNames": photoNames,
	})
}
