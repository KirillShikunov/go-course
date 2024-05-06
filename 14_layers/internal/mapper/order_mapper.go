package mapper

import (
	"14_layers/internal/dto"
	"14_layers/internal/models"
)

type OrderMapper struct {
}

func (m *OrderMapper) ConvertDTOToModel(orderDTO *dto.Order) *models.Order {
	return &models.Order{
		CustomerID: orderDTO.CustomerID,
	}
}

func (m *OrderMapper) ConvertModelsToDTOs(orders []*models.Order) []*dto.Order {
	orderDTOs := make([]*dto.Order, len(orders))

	for i, order := range orders {
		orderDTOs[i] = &dto.Order{
			ID:         order.ID,
			CustomerID: order.CustomerID,
			CreatedAt:  order.CreatedAt,
			Status:     order.Status,
			TotalPrice: order.TotalPrice,
		}
	}

	return orderDTOs
}

func NewOrderMapper() *OrderMapper {
	return &OrderMapper{}
}
