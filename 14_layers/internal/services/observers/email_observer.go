package observers

import (
	"14_layers/internal/models"
)

type EmailSender interface {
	SendEmail(userID int, mailID int)
}

type EmailObserver struct {
	sender EmailSender
}

func (o *EmailObserver) Notify(order *models.Order) {
	o.sender.SendEmail(order.UserID, order.ID)
}

func NewEmailObserver() *EmailObserver {
	return &EmailObserver{}
}
