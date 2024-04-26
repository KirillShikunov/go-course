package models

import "time"

type Order struct {
	ID        int
	Name      string
	UserId    int
	CreatedAt time.Time
}
