package main

import (
	"log"
	"net/http"

	"github.com/PacktPublishing/Go-Recipes-for-Developers/src/chp17/sort/service"
)

func main() {
	mux := service.GetServeMux()
	server := http.Server{
		Addr:    ":8088",
		Handler: mux,
	}
	log.Println(server.ListenAndServe())
}
