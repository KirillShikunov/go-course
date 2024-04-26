package repositories

import (
	"14_layers/internal/models"
	"time"
)

var orders = []*models.Order{
	{
		ID:        1,
		Name:      "Order #1",
		UserId:    1,
		CreatedAt: time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		ID:        2,
		Name:      "Order #2",
		UserId:    2,
		CreatedAt: time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		ID:        3,
		Name:      "Order #3",
		UserId:    2,
		CreatedAt: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
}

var LastID = 3

type OrderRepository struct {
}

func (r OrderRepository) Create(order *models.Order) {
	newId := LastID + 1
	order.ID = newId

	orders = append(orders, order)

	LastID = newId
}

func (r OrderRepository) List() []*models.Order {
	return orders
}
