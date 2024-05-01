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
	SentEmails []Email
}

func (m *MockEmailSender) SendEmail(userID int, mailID int) {
	m.SentEmails = append(m.SentEmails, Email{ID: mailID, UserID: userID})
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
	assert.Equal(t, expected, mockSender.SentEmails)
}
