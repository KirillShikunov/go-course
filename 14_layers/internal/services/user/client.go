package user

import (
	"14_layers/internal/cache"
	"14_layers/internal/dto"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type UserClient struct {
	cache cache.Cache
}

func (c *UserClient) GetUser(ctx context.Context, id int) (string, error) {
	key := fmt.Sprintf("user_%d", id)
	userJson, err := c.cache.Get(ctx, key)

	if err != nil {
		return "", fmt.Errorf("error getting user data: %w", err)
	}

	if userJson == "" {
		foundUser := &dto.User{ID: id, Username: "user1", Email: "test@mail.com"}
		jsonFoundUser, err := json.Marshal(foundUser)
		if err != nil {
			return "", fmt.Errorf("error marshaling user data: %w", err)
		}

		expireAt := time.Duration(5) * time.Minute
		if err := c.cache.Set(ctx, key, string(jsonFoundUser), expireAt); err != nil {
			return "", fmt.Errorf("error setting user data: %w", err)
		}
	}

	return userJson, nil
}

func NewUserClient(cache cache.Cache) *UserClient {
	return &UserClient{cache: cache}
}
