package services

import "14_layers/internal/models"

type OrderRepository interface {
	List() []*models.Order
	Create(order *models.Order)
}

type OrderObserver interface {
	Notify(order *models.Order)
}

func NewOrderManager(repository OrderRepository, observers []OrderObserver) *OrderManager {
	return &OrderManager{repository, observers}
}
