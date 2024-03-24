package console

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Manager struct {
	scanner *bufio.Scanner
}

func (c Manager) GetValue() string {
	c.scanner.Scan()
	text := c.scanner.Text()

	return strings.TrimSpace(text)
}

func (c Manager) AddBreakLine() {
	fmt.Println()
}

func NewConsoleManager() Manager {
	scanner := bufio.NewScanner(os.Stdin)
	return Manager{scanner}
}
