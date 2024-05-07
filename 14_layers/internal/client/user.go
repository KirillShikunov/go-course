package client

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
	userJson := c.cache.Get(ctx, key)

	if userJson == "" {
		foundUser := &dto.User{ID: id, Username: "user1", Email: "test@mail.com"}
		jsonFoundUser, err := json.Marshal(foundUser)
		if err != nil {
			return "", fmt.Errorf("error marshaling user data: %w", err)
		}

		expireAt := time.Duration(5) * time.Minute
		c.cache.Set(ctx, key, string(jsonFoundUser), expireAt)
	}

	return userJson, nil
}

func NewUserClient(cache cache.Cache) *UserClient {
	return &UserClient{cache: cache}
}
