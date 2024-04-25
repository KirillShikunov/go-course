package models

import "time"

type Order struct {
	ID        int
	Name      string
	CreatedAt time.Time
}
