package main

import (
	"log"
	"net/http"

	"github.com/bserdar/go-recipes-book/chp15/sort/service"
)

func main() {
	mux := service.GetServeMux()
	server := http.Server{
		Addr:    ":8088",
		Handler: mux,
	}
	log.Println(server.ListenAndServe())
}
