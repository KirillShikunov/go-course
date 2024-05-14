package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func NewConnection(dsn string) *gorm.DB {
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return connection
}
