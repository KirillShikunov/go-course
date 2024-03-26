package main

import (
	"06_interface/booking"
	"06_interface/booking/dto"
	"06_interface/booking/step"
	"06_interface/console"
	"06_interface/tour"
	"06_interface/transport"
	"fmt"
)

func main() {
	steps := generateSteps()
	bookProcessor := booking.NewBookingProcessor(steps)

	dtoProcessor := dto.ProcessorDTO{}
	consoleManager := console.NewConsoleManager()

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

func generateSteps() []step.Step {
	var steps []step.Step

	steps = append(steps, step.NewBudgetStep())

	tourManager := generateTourManager()
	steps = append(steps, step.NewTourStep(tourManager))

	steps = append(steps, step.NewClientStep())

	return steps
}

func generateTourManager() tour.Manager {
	var tours []tour.Tour

	car := transport.NewCar("456DFG")
	plane := transport.NewPlane("Airbus A320")
	boat := transport.NewBoat("Лотка з ухилятнами")
	train := transport.NewTrain("Eurostar")

	tours = append(tours, tour.NewTour(1, "Вроцлав - Відень", 500, []tour.Transport{car, plane}))
	tours = append(tours, tour.NewTour(2, "Прага - Берлін", 400, []tour.Transport{car, train}))
	tours = append(tours, tour.NewTour(3, "Київ - Угорщина(через Тису)", 3000, []tour.Transport{car, boat}))

	return tour.NewTourManager(tours)
}

func handleError(err error) bool {
	if err != nil {
		fmt.Println(err)
		return true
	}
	return false
}
