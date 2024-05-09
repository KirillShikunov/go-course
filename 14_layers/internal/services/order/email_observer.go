package order

import (
	"14_layers/internal/dto"
	"14_layers/internal/models"
	"context"
	"fmt"
)

const OrderCreatedMailID = 123

type EmailSender interface {
	SendEmail(email string, mailID int)
}

type UserService interface {
	GetUser(ctx context.Context, id int) (*dto.User, error)
}

type EmailObserver struct {
	sender  EmailSender
	service UserService
}

func (o *EmailObserver) Notify(ctx context.Context, order *models.Order) {
	user, err := o.service.GetUser(ctx, order.CustomerID)
	if err != nil {
		fmt.Println(err) // log error
		return
	}

	o.sender.SendEmail(user.Email, OrderCreatedMailID)
}

func NewEmailObserver(sender EmailSender, service UserService) *EmailObserver {
	return &EmailObserver{sender, service}
}
