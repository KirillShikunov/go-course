package repositories

import (
	"14_layers/internal/models"
	"time"
)

var orders = []*models.Order{
	{
		ID:        1,
		Name:      "Order #1",
		UserID:    1,
		CreatedAt: time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		ID:        2,
		Name:      "Order #2",
		UserID:    2,
		CreatedAt: time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		ID:        3,
		Name:      "Order #3",
		UserID:    2,
		CreatedAt: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
}

var LastID = 3

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
