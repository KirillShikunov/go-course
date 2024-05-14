package services

import (
	"21_uber-fx/internal/models"
	"context"
	"math/rand"
)

type Repository interface {
	List(ctx context.Context) ([]*models.Order, error)
	Create(ctx context.Context, order *models.Order) error
}

func NewOrderManager(repository Repository) *Manager {
	return &Manager{repository}
}

type Manager struct {
	repository Repository
}

func (m *Manager) List(ctx context.Context) ([]*models.Order, error) {
	return m.repository.List(ctx)
}

func (m *Manager) Create(ctx context.Context, order *models.Order) error {
	order.Status = models.OrderStatusCreated
	order.TotalPrice = rand.Intn(100)
	err := m.repository.Create(ctx, order)

	return err
}
