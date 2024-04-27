package main

import (
	"14_layers/internal/api"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	api.RegisterRoutes(router)

	log.Println("Starting web server on localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}
