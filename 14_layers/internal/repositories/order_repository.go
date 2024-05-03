package repositories

import (
	"14_layers/internal/models"
	"gorm.io/gorm"
)

func NewOrderRepository(connection *gorm.DB) *OrderRepository {
	return &OrderRepository{connection}
}

type OrderRepository struct {
	connection *gorm.DB
}

func (r *OrderRepository) Create(order *models.Order) {
	r.connection.Create(&order)
}

func (r *OrderRepository) List() []*models.Order {
	var orders []*models.Order
	r.connection.Find(&orders)

	return orders
}
