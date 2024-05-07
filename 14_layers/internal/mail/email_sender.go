package mail

import "fmt"

type EmailSender struct {
}

func (e *EmailSender) SendEmail(email string, mailID int) {
	fmt.Printf("EmailSender: Email: %s; Mail ID: %d", email, mailID)
}

func NewEmailSender() *EmailSender {
	return &EmailSender{}
}
