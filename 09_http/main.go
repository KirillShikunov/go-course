package main

import (
	"09_http/api"
	"log"
	"net/http"
)

func main() {
	api.Init()

	log.Println("Server is running. Logging page http://localhost:8020/login")
	log.Fatal(http.ListenAndServe(":8020", nil))
}
