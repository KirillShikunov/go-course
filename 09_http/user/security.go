package user

import (
	"crypto/sha1"
	"fmt"
)

var users = map[int]string{
	1: hashPassword("1"),
	2: hashPassword("2"),
}

func hashPassword(password string) string {
	hasher := sha1.New()
	hasher.Write([]byte(password))
	return fmt.Sprintf("%x", hasher.Sum(nil))
}

func CheckAccess(id int, password string) bool {
	hashedPassword, exists := users[id]
	if !exists {
		return false
	}
	return hashPassword(password) == hashedPassword
}
