package main

import (
	"log"

	"github.com/mockallthethings/photo-upload-example/handler"
)

func main() {
	err := handler.Serve()
	if err != nil {
		log.Fatal(err)
	}
	select {}
}
