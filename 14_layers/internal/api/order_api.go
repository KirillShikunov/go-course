package api

import (
	"14_layers/internal/dto"
	"14_layers/internal/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type OrderManager interface {
	List() []*models.Order
	Create(orderDTO *models.Order) error
}

func NewOrderAPI(manager OrderManager) *OrderAPI {
	return &OrderAPI{manager}
}

type OrderAPI struct {
	manager OrderManager
}

func (api *OrderAPI) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/orders", api.listOrders).Methods("GET")
	router.HandleFunc("/orders", api.createOrder).Methods("POST")
}

func (api *OrderAPI) listOrders(w http.ResponseWriter, r *http.Request) {
	orders := api.manager.List()
	orderDTOs := api.convertModelsToDTOs(orders)

	err := json.NewEncoder(w).Encode(orderDTOs)
	if err != nil {
		http.Error(w, "Failed to encode orders", http.StatusInternalServerError)
	}
}

func (api *OrderAPI) createOrder(w http.ResponseWriter, r *http.Request) {
	var orderDTO *dto.Order

	if err := json.NewDecoder(r.Body).Decode(&orderDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := api.convertDTOToModel(orderDTO)

	if err := api.manager.Create(order); err != nil {
		http.Error(w, fmt.Sprintf("Failed to create order: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (api *OrderAPI) convertDTOToModel(orderDTO *dto.Order) *models.Order {
	return &models.Order{
		Name:   orderDTO.Name,
		UserID: orderDTO.UserID,
	}
}

func (api *OrderAPI) convertModelsToDTOs(orders []*models.Order) []*dto.Order {
	orderDTOs := make([]*dto.Order, len(orders))

	for i, order := range orders {
		orderDTOs[i] = &dto.Order{
			ID:     order.ID,
			Name:   order.Name,
			UserID: order.UserID,
		}
	}

	return orderDTOs
}
