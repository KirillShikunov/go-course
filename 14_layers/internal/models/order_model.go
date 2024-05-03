package models

import "time"

const OrderStatusCreated = 1

type Order struct {
	ID         int       `json:"id"`
	Status     int       `json:"status"`
	TotalPrice int       `json:"total_price"`
	CustomerID int       `json:"customer_id"`
	CreatedAt  time.Time `json:"created_at"`
}
