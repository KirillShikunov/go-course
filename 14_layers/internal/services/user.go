package services

import (
	"14_layers/internal/client"
	"14_layers/internal/dto"
	"context"
	"encoding/json"
	"fmt"
)

type UserService struct {
	client *client.UserClient
}

func (s *UserService) GetUser(ctx context.Context, id int) (*dto.User, error) {
	data, err := s.client.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return s.unmarshalUser(data)
}

func (s *UserService) unmarshalUser(data string) (*dto.User, error) {
	var user dto.User
	err := json.Unmarshal([]byte(data), &user)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling user data: %w", err)
	}
	return &user, nil
}

func NewUserService(client *client.UserClient) *UserService {
	return &UserService{client}
}
