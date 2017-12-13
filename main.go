package main

import (
	"log"

	"github.com/mockallthethings/photo-upload-example/db"
	"github.com/mockallthethings/photo-upload-example/handler"
)

func main() {
	db, err := db.OpenDB()
	if err != nil {
		log.Panicf("Could not start - error connecting to DB: %v", err)
	}

	h, err := handler.New(db)
	h.Run()
}
