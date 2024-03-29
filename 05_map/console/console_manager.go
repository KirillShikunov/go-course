package console

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ConsoleManager struct {
	scanner *bufio.Scanner
}

func (consoleManager ConsoleManager) ShowRows(rows []string) {
	fmt.Println("Результат пошуку:")

	if len(rows) == 0 {
		fmt.Printf("Не вдалося знайти рядкі які б відповідали крітерію пошуку. \n")
		return
	}

	for i, row := range rows {
		fmt.Printf("#%d %s \n", i+1, row)
	}
}

func (consoleManager ConsoleManager) ShowAvgSearchDuration(totalSearchDuration int64, countSearch int) {
	fmt.Printf("Середній час пошуку: %dнаносекунд\n", totalSearchDuration/int64(countSearch))
}

func (consoleManager ConsoleManager) GetSearchWord() string {
	consoleManager.scanner.Scan()
	searchWord := consoleManager.scanner.Text()

	return strings.TrimSpace(searchWord)
}

func (consoleManager ConsoleManager) AddBreakLine() {
	fmt.Println()
}

func NewConsoleManager() ConsoleManager {
	scanner := bufio.NewScanner(os.Stdin)
	return ConsoleManager{scanner}
}
