package services

import "14_layers/internal/models"

type OrderRepository interface {
	List() []*models.Order
	Create(order *models.Order)
}

func NewOrderManager(repository OrderRepository) *OrderManager {
	return &OrderManager{repository}
}
