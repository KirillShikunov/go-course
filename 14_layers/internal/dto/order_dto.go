package dto

import (
	"time"
)

type Order struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
