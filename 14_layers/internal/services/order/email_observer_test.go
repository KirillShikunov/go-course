package order

import (
	"14_layers/internal/dto"
	"14_layers/internal/models"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type Email struct {
	ID    int
	Email string
}

type MockEmailSender struct {
	SentEmails []Email
}

func (m *MockEmailSender) SendEmail(email string, mailID int) {
	m.SentEmails = append(m.SentEmails, Email{ID: mailID, Email: email})
}

type MockUserService struct {
}

func (m *MockUserService) GetUser(ctx context.Context, id int) (*dto.User, error) {
	return &dto.User{ID: id, Username: "John", Email: "john@gmail.com"}, nil
}

func TestNotify(t *testing.T) {
	order := &models.Order{
		ID:         456,
		CustomerID: 123,
		TotalPrice: 44,
		Status:     OrderCreatedMailID,
		CreatedAt:  time.Now(),
	}

	mockSender := &MockEmailSender{}
	userService := &MockUserService{}

	ctx := context.Background()

	observer := &EmailObserver{sender: mockSender, service: userService}
	observer.Notify(ctx, order)

	expected := []Email{{ID: order.Status, Email: "john@gmail.com"}}
	assert.Equal(t, expected, mockSender.SentEmails)
}
