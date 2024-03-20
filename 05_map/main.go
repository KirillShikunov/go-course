package main

import (
	"05_map/benchmark"
	"05_map/console"
	"05_map/file"
	"05_map/text"
	"fmt"
)

func main() {
	rows, err := file.GetRowsFromFile("texts.txt")
	if err != nil {
		fmt.Println("Не вдалося прочитати рядки з файлу: ", err)
		return
	}

	textManager := text.NewTextManager(rows)

	consoleManager := console.NewConsoleManager()
	timeTracker := benchmark.TimeTracker{}

	var matchedRows []string
	var duration string

	for {
		fmt.Println("Введить слово пошуку:")

		searchWord := consoleManager.GetSearchWord()

		duration = timeTracker.Track(func() {
			matchedRows = textManager.SearchWordByIndex(searchWord)
		})

		fmt.Printf("Час пошуку по індексу: %s\n", duration)
		consoleManager.ShowRows(matchedRows)

		fmt.Println() // Break row, add just for separate result two search functions

		duration = timeTracker.Track(func() {
			matchedRows = textManager.SearchWordByContains(searchWord)
		})

		fmt.Printf("Час пошуку з використанням стандартник функцій: %s\n", duration)
		consoleManager.ShowRows(matchedRows)
	}
}
