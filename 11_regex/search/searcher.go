package search

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func FindInFile(fileName string, regex *regexp.Regexp) ([]string, error) {
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

	scanner := bufio.NewScanner(file)
	var result []string
	for scanner.Scan() {
		line := scanner.Text()

		matches := regex.FindAllString(line, -1)
		for _, match := range matches {
			result = append(result, match)
		}
	}

	return result, nil
}
