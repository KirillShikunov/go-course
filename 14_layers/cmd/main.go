package main

import (
	"14_layers/internal/app"
	"14_layers/internal/di"
	"github.com/gorilla/mux"
)

func main() {
	container := di.NewServiceContainer()
	var newApp = app.NewApp(container.Load())

	router := mux.NewRouter()
	newApp.RunServer(router)
}
