package main

import (
	"flag"
	"log"

	"github.com/mockallthethings/photo-upload-example/handler"
)

func main() {
	portPtr := flag.Int("port", 0, "port on which to listen")
	flag.Parse()
	if *portPtr == 0 {
		flag.Usage()
		return
	}

	err := handler.Serve(*portPtr)
	if err != nil {
		log.Fatal(err)
	}
	select {}
}
