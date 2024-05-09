package order

import (
	"14_layers/internal/models"
	"context"
	"math/rand"
)

type Repository interface {
	List(ctx context.Context) ([]*models.Order, error)
	Create(ctx context.Context, order *models.Order) error
}

type Observer interface {
	Notify(ctx context.Context, order *models.Order)
}

func NewOrderManager(repository Repository, observers []Observer) *Manager {
	return &Manager{repository, observers}
}

type Manager struct {
	repository Repository
	observers  []Observer
}

func (m *Manager) List(ctx context.Context) ([]*models.Order, error) {
	return m.repository.List(ctx)
}

func (m *Manager) Create(ctx context.Context, order *models.Order) error {
	order.Status = models.OrderStatusCreated
	order.TotalPrice = rand.Intn(100)
	err := m.repository.Create(ctx, order)

	m.notifyObservers(ctx, order)

	return err
}

func (m *Manager) notifyObservers(ctx context.Context, order *models.Order) {
	for _, observer := range m.observers {
		observer.Notify(ctx, order)
	}
}
