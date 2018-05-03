package main

import (
	"flag"
	"log"

	"github.com/mockallthethings/photo-upload-example/handler"
)

func main() {
	portPtr := flag.Int("port", 80, "port on which to listen")
	flag.Parse()

	err := handler.Serve(*portPtr)
	if err != nil {
		log.Fatal(err)
	}
	select {}
}
