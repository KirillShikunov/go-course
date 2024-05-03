package repositories

import (
	"14_layers/internal/models"
	"time"
)

var orders []*models.Order

var LastID = 0

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{}
}

type OrderRepository struct {
}

func (r OrderRepository) Create(order *models.Order) {
	newID := LastID + 1

	order.ID = newID
	order.CreatedAt = time.Now()

	orders = append(orders, order)

	LastID = newID
}

func (r OrderRepository) List() []*models.Order {
	return orders
}
