package api

import (
	"14_layers/internal/dto"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type OrderApi struct {
	manager OrderManager
}

func (api *OrderApi) listOrders(w http.ResponseWriter, r *http.Request) {
	orders := api.manager.List()
	err := json.NewEncoder(w).Encode(orders)
	if err != nil {
		http.Error(w, "Failed to encode orders", http.StatusInternalServerError)
	}
}

func (api *OrderApi) createOrder(w http.ResponseWriter, r *http.Request) {
	var orderDTO *dto.OrderDTO

	if err := json.NewDecoder(r.Body).Decode(&orderDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := api.manager.Create(orderDTO); err != nil {
		http.Error(w, fmt.Sprintf("Failed to create order: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (api *OrderApi) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/orders", api.listOrders).Methods("GET")
	router.HandleFunc("/orders", api.createOrder).Methods("POST")
}
