package handler

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mockallthethings/photo-upload-example/db"
	"github.com/mockallthethings/photo-upload-example/s3"
)

type handler struct {
	db      *sql.DB
	S3      s3.Client
	loading bool
}

func (h *handler) loadingMiddleware(c *gin.Context) {
	if h.loading {
		c.HTML(http.StatusOK, "loading.tmpl", nil)
		c.Abort()
	}
}

func Serve() error {
	// show a loading template while we wait for DB to load
	r := gin.Default()
	h := &handler{
		loading: true,
	}
	r.Use(h.loadingMiddleware)

	r.LoadHTMLGlob("/public/*.tmpl")
	r.StaticFile("/styles/index.css", "/public/styles/index.css")

	r.GET("/", h.indexHandler)
	r.POST("/albums/create", h.createAlbumHandler)
	r.POST("/album/:albumID/upload", h.uploadHandler)
	r.GET("/album/:albumID/view", h.indexHandler)

	// Start serving in background goroutine immediately
	// so that we can show 'loading' while DB connection loads
	go func() {
		log.Fatal(r.Run(":80"))
	}()

	// wait for DB connection (or die if it takes too long)
	db, err := db.OpenDB()
	if err != nil {
		return err
	}
	h.db = db

	// Set up S3 uploader
	h.S3, err = s3.NewClient()
	if err != nil {
		return err
	}

	h.loading = false
	return nil
}
