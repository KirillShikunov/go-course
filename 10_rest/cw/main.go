package main

import (
	"10_rest/cw/user"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/users", GetUsers).Methods(http.MethodGet)
	r.HandleFunc("/api/users/{id}", GetUser).Methods(http.MethodGet)

	fmt.Println("Server is running on port 8020")
	log.Fatal(http.ListenAndServe("localhost:8020", r))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	foundUser, err := user.Get(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	if err = json.NewEncoder(w).Encode(foundUser); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := user.List()
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
