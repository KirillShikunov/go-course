package auth

import (
	"encoding/base64"
	"strconv"
)

func GenerateToken(id int) string {
	token := base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(id)))
	return token
}

func DecodeToken(token string) (int, bool) {
	decodedBytes, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return 0, false
	}

	id, err := strconv.Atoi(string(decodedBytes))
	if err != nil {
		return 0, false
	}

	return id, true
}
