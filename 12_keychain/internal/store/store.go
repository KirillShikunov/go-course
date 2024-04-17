package store

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

const filePath = "passwords.txt"

func SavePassword(name, password string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintf("%s:%s\n", name, password))
	return err
}

func GetPassword(name string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if parts[0] == name {
			return parts[1], nil
		}
	}

	return "", errors.New("password not found")
}

func ListPasswords() []string {
	var names []string
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return names
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) > 0 {
			names = append(names, parts[0])
		}
	}

	return names
}
