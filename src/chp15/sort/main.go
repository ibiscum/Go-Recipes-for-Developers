package main

import (
	"log"
	"net/http"

	"github.com/PacktPublishing/GO-Cookbook-Top-Techniques/src/chp15/sort/service"
)

func main() {
	mux := service.GetServeMux()
	server := http.Server{
		Addr:    ":8088",
		Handler: mux,
	}
	log.Println(server.ListenAndServe())
}
