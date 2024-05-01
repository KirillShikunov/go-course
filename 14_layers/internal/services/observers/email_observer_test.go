package observers

import (
	"14_layers/internal/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Email struct {
	ID     int
	UserID int
}

type MockEmailSender struct {
	sentEmails []Email
}

func (m *MockEmailSender) SendEmail(userID int, mailID int) {
	m.sentEmails = append(m.sentEmails, Email{ID: mailID, UserID: userID})
}

func (m *MockEmailSender) GetSentEmails() []Email {
	return m.sentEmails
}

func TestNotify(t *testing.T) {
	order := &models.Order{
		ID:     456,
		UserID: 123,
	}

	mockSender := &MockEmailSender{}
	observer := &EmailObserver{sender: mockSender}
	observer.Notify(order)

	expected := []Email{{ID: OrderCreatedMailID, UserID: order.UserID}}
	assert.Equal(t, expected, mockSender.GetSentEmails())
}
