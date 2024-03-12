package main

import (
	"fmt"
)

func main() {
	textManager, err := getTextManager()
	if err != nil {
		fmt.Println("Помилка при ініцілазації текст менеджера: ", err)
		return
	}

	consoleManager := getConsoleManager()

	for {
		fmt.Println("Введить слово пошуку:")

		searchWord := consoleManager.getSearchWord()
		matchedRows := textManager.searchRows(searchWord)

		consoleManager.showRows(matchedRows)
	}
}
