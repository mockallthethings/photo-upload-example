package handler

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type handler struct {
	db *sql.DB
}

type Handler interface {
	Run()
}

func (h *handler) Run() {
	r := gin.Default()
	r.LoadHTMLGlob("public/*.tmpl")
	r.GET("/", h.indexHandler)
	r.POST("/upload", h.uploadHandler)
	r.GET("/view", h.viewHandler)
	r.Run(":80")
}

func New(db *sql.DB) (Handler, error) {
	return &handler{
		db: db,
	}, nil
}
