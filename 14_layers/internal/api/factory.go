package api

import "14_layers/internal/dto"

type OrderManager interface {
	List() []*dto.OrderDTO
	Create(orderDTO *dto.OrderDTO) error
}

func NewOrderApi(manager OrderManager) *OrderApi {
	return &OrderApi{manager}
}
