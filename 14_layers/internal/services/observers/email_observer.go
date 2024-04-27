package observers

import (
	"14_layers/internal/models"
	"fmt"
)

type EmailObserver struct {
}

func (o *EmailObserver) Notify(order *models.Order) {
	fmt.Printf("EmailObserver: Sending email for user %d; Order ID: %d", order.UserID, order.ID)
}

func NewEmailObserver() *EmailObserver {
	return &EmailObserver{}
}
