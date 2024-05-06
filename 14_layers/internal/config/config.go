package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading env files:" + err.Error())
	}
}

func Env(key string) string {
	return os.Getenv(key)
}
