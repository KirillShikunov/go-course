package observers

import (
	"14_layers/internal/models"
)

const OrderCreatedMailID = 123

type EmailSender interface {
	SendEmail(userID int, mailID int)
}

type EmailObserver struct {
	sender EmailSender
}

func (o *EmailObserver) Notify(order *models.Order) {
	o.sender.SendEmail(order.UserID, OrderCreatedMailID)
}

func NewEmailObserver() *EmailObserver {
	return &EmailObserver{}
}
