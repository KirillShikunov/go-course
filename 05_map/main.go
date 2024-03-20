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
	var duration int64
	var totalSearchDurationByIndex int64
	var totalSearchDurationByContain int64

	countSearch := 1

	for {
		fmt.Println("Введить слово пошуку:")

		searchWord := consoleManager.GetSearchWord()

		duration = timeTracker.Track(func() {
			matchedRows = textManager.SearchWordByIndex(searchWord)
		})
		totalSearchDurationByIndex += duration

		fmt.Printf("Час пошуку по індексу: %dнаносекунд\n", duration)
		consoleManager.ShowAvgSearchDuration(totalSearchDurationByIndex, countSearch)
		consoleManager.ShowRows(matchedRows)

		consoleManager.AddBreakLine()

		duration = timeTracker.Track(func() {
			matchedRows = textManager.SearchWordByContains(searchWord)
		})

		totalSearchDurationByContain += duration

		fmt.Printf("Час пошуку з використанням стандартник функцій: %dнаносекунд\n", duration)
		consoleManager.ShowAvgSearchDuration(totalSearchDurationByContain, countSearch)
		consoleManager.ShowRows(matchedRows)

		consoleManager.AddBreakLine()
		countSearch++
	}
}
