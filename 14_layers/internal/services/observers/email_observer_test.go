package observers

import (
	"14_layers/internal/models"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockEmailSender struct {
	mock.Mock
}

func (m *MockEmailSender) SendEmail(userID int, mailID int) {
	m.Called(userID, mailID)
}

func TestNotify(t *testing.T) {
	mockSender := &MockEmailSender{}
	observer := &EmailObserver{sender: mockSender}

	order := &models.Order{
		ID:     456,
		UserID: 123,
	}

	mockSender.On("SendEmail", order.UserID, order.ID).Return(nil)

	observer.Notify(order)
}
