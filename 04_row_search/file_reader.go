package main

import (
	"bufio"
	"fmt"
	"os"
)

type FileReader struct {
}

func (fileReader FileReader) read(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("помилка при відкритті файлу: %w", err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Не вдалося закрити файл:", err)
		}
	}(file)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return lines, fmt.Errorf("помилка при читанні файлу: %w", err)
	}

	return lines, nil
}
