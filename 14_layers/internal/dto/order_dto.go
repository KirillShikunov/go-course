package dto

import "time"

type OrderDTO struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}
