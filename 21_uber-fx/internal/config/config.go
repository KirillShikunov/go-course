package config

import (
	"os"
)

func Env(key string) string {
	return os.Getenv(key)
}

type DBConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

func NewDBConfig() *DBConfig {
	return &DBConfig{
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		"disable",
	}
}
