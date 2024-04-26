package services

import (
	"14_layers/internal/dto"
	"14_layers/internal/models"
	"time"
)

type OrderManager struct {
	repository OrderRepository
	observers  []OrderObserver
}

func (m *OrderManager) List() []*dto.OrderDTO {
	var DTOs []*dto.OrderDTO

	for _, order := range m.repository.List() {
		DTOs = append(DTOs, dto.NewOrderFromModel(order))
	}

	return DTOs
}

func (m *OrderManager) Create(orderDTO *dto.OrderDTO) error {
	order := &models.Order{
		ID:        orderDTO.ID,
		Name:      orderDTO.Name,
		UserId:    orderDTO.UserId,
		CreatedAt: time.Now(),
	}

	m.repository.Create(order)

	m.notifyObservers(order)

	return nil
}

func (m *OrderManager) notifyObservers(order *models.Order) {
	for _, observer := range m.observers {
		observer.Notify(order)
	}
}
