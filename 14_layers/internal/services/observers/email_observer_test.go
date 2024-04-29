package observers

import (
	"14_layers/internal/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockEmailSender struct {
	order   *models.Order
	testing *testing.T
}

func (m *MockEmailSender) SendEmail(userID int, mailID int) {
	assert.Equal(m.testing, m.order.UserID, userID)
	assert.Equal(m.testing, OrderCreatedMailID, mailID)
}

func TestNotify(t *testing.T) {
	order := &models.Order{
		ID:     456,
		UserID: 123,
	}

	mockSender := &MockEmailSender{order, t}
	observer := &EmailObserver{sender: mockSender}

	observer.Notify(order)
}
