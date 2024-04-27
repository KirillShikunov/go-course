package models

import "time"

type Order struct {
	ID        int
	Name      string
	UserID    int
	CreatedAt time.Time
}
