package dto

import (
	"14_layers/internal/models"
	"time"
)

type OrderDTO struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func NewOrderFromModel(order *models.Order) *OrderDTO {
	return &OrderDTO{
		ID:        order.ID,
		Name:      order.Name,
		UserId:    order.UserId,
		CreatedAt: order.CreatedAt,
	}
}
