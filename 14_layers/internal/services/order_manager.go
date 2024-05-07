package services

import (
	"14_layers/internal/models"
	"context"
	"math/rand"
)

type OrderRepository interface {
	List(ctx context.Context) ([]*models.Order, error)
	Create(ctx context.Context, order *models.Order) error
}

type OrderObserver interface {
	Notify(ctx context.Context, order *models.Order)
}

func NewOrderManager(repository OrderRepository, observers []OrderObserver) *OrderManager {
	return &OrderManager{repository, observers}
}

type OrderManager struct {
	repository OrderRepository
	observers  []OrderObserver
}

func (m *OrderManager) List(ctx context.Context) ([]*models.Order, error) {
	return m.repository.List(ctx)
}

func (m *OrderManager) Create(ctx context.Context, order *models.Order) error {
	order.Status = models.OrderStatusCreated
	order.TotalPrice = rand.Intn(100)
	err := m.repository.Create(ctx, order)

	m.notifyObservers(ctx, order)

	return err
}

func (m *OrderManager) notifyObservers(ctx context.Context, order *models.Order) {
	for _, observer := range m.observers {
		observer.Notify(ctx, order)
	}
}
