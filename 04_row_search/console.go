package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ConsoleManager struct {
	scanner *bufio.Scanner
}

func (consoleManager ConsoleManager) showRows(rows []string) {
	if len(rows) == 0 {
		fmt.Printf("Не вдалося знайти рядкі які б відповідали крітерію пошуку. \n")
		return
	}

	fmt.Println("Результат пошуку:")
	for i, row := range rows {
		fmt.Printf("#%d %s \n", i+1, row)
	}
}

func (consoleManager ConsoleManager) getSearchWord() string {
	consoleManager.scanner.Scan()
	searchWord := consoleManager.scanner.Text()

	return strings.TrimSpace(searchWord)
}

func getConsoleManager() ConsoleManager {
	scanner := bufio.NewScanner(os.Stdin)
	return ConsoleManager{scanner}
}
