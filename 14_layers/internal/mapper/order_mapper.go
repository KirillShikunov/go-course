package mapper

import (
	"14_layers/internal/dto"
	"14_layers/internal/models"
)

type OrderMapper struct {
}

func (m *OrderMapper) ConvertDTOToModel(orderDTO *dto.Order) *models.Order {
	return &models.Order{
		ID:     orderDTO.ID,
		Name:   orderDTO.Name,
		UserID: orderDTO.UserID,
	}
}

func (m *OrderMapper) ConvertModelsToDTOs(orders []*models.Order) []*dto.Order {
	orderDTOs := make([]*dto.Order, len(orders))

	for i, order := range orders {
		orderDTOs[i] = &dto.Order{
			ID:     order.ID,
			Name:   order.Name,
			UserID: order.UserID,
		}
	}

	return orderDTOs
}

func NewOrderMapper() OrderMapper {
	return OrderMapper{}
}
