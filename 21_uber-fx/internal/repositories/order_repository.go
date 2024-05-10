package repositories

import (
	"21_uber-fx/internal/models"
	"context"
	"gorm.io/gorm"
)

func NewOrderRepository(connection *gorm.DB) *OrderRepository {
	return &OrderRepository{connection}
}

type OrderRepository struct {
	connection *gorm.DB
}

func (r *OrderRepository) Create(ctx context.Context, order *models.Order) error {
	return r.connection.WithContext(ctx).Create(&order).Error
}

func (r *OrderRepository) List(ctx context.Context) ([]*models.Order, error) {
	var orders []*models.Order

	err := r.connection.WithContext(ctx).Find(&orders).Error

	return orders, err
}
