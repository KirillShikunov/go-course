package main

import (
	"14_layers/internal/api"
	"14_layers/internal/config"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	api.RegisterRoutes(router)

	port := config.Env("PORT")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
