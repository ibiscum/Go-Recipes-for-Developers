package main

import (
	"log"
	"net/http"

	"github.com/ibiscum/Go-Recipes-for-Developers/chp17/sorting/service"
)

func main() {
	mux := service.GetServeMux()
	server := http.Server{
		Addr:    ":8088",
		Handler: mux,
	}
	log.Println(server.ListenAndServe())
}
