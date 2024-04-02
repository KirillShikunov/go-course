package main

import (
	"10_rest/hw/api"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	api.InitRoutes(r)

	fmt.Println("Server is running localhost:8020")
	log.Fatal(http.ListenAndServe("localhost:8020", r))
}
