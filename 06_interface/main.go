package main

import (
	"06_interface/booking"
	"06_interface/booking/dto"
	"06_interface/console"
	"fmt"
)

func main() {
	consoleManager := console.NewConsoleManager()
	bookProcessor := booking.BuildBookProcessor()

	dtoProcessor := dto.ProcessorDTO{}

	for {
		bookStep, ok := bookProcessor.GetStep()
		if handleError(ok) {
			return
		}

		ok = bookStep.Print(&dtoProcessor)
		if handleError(ok) {
			return
		}

		value := consoleManager.GetValue()

		ok = bookStep.Process(value, &dtoProcessor)
		if handleError(ok) {
			return
		}

		if bookProcessor.IsLastStep() {
			consoleManager.AddBreakLine()
			fmt.Println("Вітаю! Ви завершили процес бронювання.")
			dtoProcessor.Print()
			return
		}

		bookProcessor.IncrementStep()
	}
}

func handleError(err error) bool {
	if err != nil {
		fmt.Println(err)
		return true
	}
	return false
}
