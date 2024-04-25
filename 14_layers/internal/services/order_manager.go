package services

import (
	"14_layers/internal/dto"
	"14_layers/internal/models"
	"time"
)

type OrderManager struct {
	repository OrderRepository
}

func (m *OrderManager) List() []*dto.OrderDTO {
	var DTOs []*dto.OrderDTO

	for _, order := range m.repository.List() {
		DTOs = append(DTOs, &dto.OrderDTO{
			ID:        order.ID,
			Name:      order.Name,
			CreatedAt: order.CreatedAt,
		})
	}

	return DTOs
}

func (m *OrderManager) Create(orderDTO *dto.OrderDTO) error {
	order := &models.Order{
		ID:        orderDTO.ID,
		Name:      orderDTO.Name,
		CreatedAt: time.Now(),
	}

	m.repository.Create(order)

	return nil
}
