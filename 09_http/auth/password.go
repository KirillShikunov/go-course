package auth

import (
	"golang.org/x/crypto/bcrypt"
)

var users = map[int]string{
	1: hashPassword("1"),
	2: hashPassword("2"),
}

func CheckAccess(id int, password string) bool {
	hashedPassword, exists := users[id]
	if !exists {
		return false
	}
	return checkPasswordHash(password, hashedPassword)
}

func hashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
