package main

import (
	"14_layers/internal/app"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	a := app.NewApp(
		app.WithAPIRoutes(router),
	)

	a.Run(router)
}
