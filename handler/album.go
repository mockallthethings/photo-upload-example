package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *handler) createAlbumHandler(c *gin.Context) {
	albumName := c.PostForm("albumname")
	if albumName == "" {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("Album name is required"))
		return
	}

	albumID, err := uuid.NewRandom()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	stmt, err := h.db.Prepare(`
INSERT INTO photos.dim_album (
  id,
	name
) VALUES (?, ?);`)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	_, err = stmt.Exec(
		albumID,
		albumName,
	)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Use PRG pattern: https://en.wikipedia.org/wiki/Post/Redirect/Get
	c.Redirect(http.StatusFound, "/album/"+albumID.String()+"/view")
}
