package user

import (
	"14_layers/internal/dto"
	"context"
	"encoding/json"
	"fmt"
)

type Service struct {
	client *UserClient
}

func (s *Service) GetUser(ctx context.Context, id int) (*dto.User, error) {
	data, err := s.client.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return s.unmarshalUser(data)
}

func (s *Service) unmarshalUser(data string) (*dto.User, error) {
	var user dto.User
	err := json.Unmarshal([]byte(data), &user)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling user data: %w", err)
	}
	return &user, nil
}

func NewUserService(client *UserClient) *Service {
	return &Service{client}
}
