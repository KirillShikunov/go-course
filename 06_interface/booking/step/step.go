package step

import (
	"06_interface/booking/dto"
	"06_interface/client"
	"06_interface/tour"
	"fmt"
	"strconv"
)

type Step interface {
	Print(dto *dto.ProcessorDTO) error
	Process(value string, dto *dto.ProcessorDTO) error
}

type BudgetStep struct {
}

func (b BudgetStep) Print(dto *dto.ProcessorDTO) error {
	fmt.Println("Вкажіть ваш бюджет:")

	return nil
}

func (b BudgetStep) Process(value string, dto *dto.ProcessorDTO) error {
	budget, ok := strconv.Atoi(value)
	if ok != nil {
		return fmt.Errorf("бюджет повинен бути цілим числом")
	}

	dto.SetBudget(budget)

	return nil
}

type TourStep struct {
	manager tour.Manager
}

func (t TourStep) Print(dto *dto.ProcessorDTO) error {
	fmt.Println("Тури за вашим бюджетом:")
	toursByBudget := t.manager.FindByBudget(dto.GetBudget())
	if len(toursByBudget) == 0 {
		return fmt.Errorf("турів за вашим бюджетом не знайдено")
	}

	for _, tourItem := range toursByBudget {
		fmt.Printf("#%d %s | Ціна: %dгрн\n", tourItem.GetId(), tourItem.GetName(), tourItem.GetPrice())
	}

	return nil
}

func (t TourStep) Process(value string, dto *dto.ProcessorDTO) error {
	tourId, ok := strconv.Atoi(value)
	if ok != nil {
		return fmt.Errorf("індекс повинен бути цілим числом")
	}

	tourChoice, ok := t.manager.Get(tourId)
	if ok != nil {
		return ok
	}

	dto.SetTour(tourChoice)

	return nil
}

type ClientNameStep struct {
}

func (n ClientNameStep) Print(dto *dto.ProcessorDTO) error {
	fmt.Println("Вкажіть ваше ім'я:")
	return nil
}

func (n ClientNameStep) Process(value string, dto *dto.ProcessorDTO) error {
	newClient := client.NewClient(value)
	dto.SetClient(&newClient)
	return nil
}

func BuildBookingSteps() []Step {
	var steps []Step

	steps = append(steps, BudgetStep{})

	tourManager := tour.BuildTourManager()
	steps = append(steps, TourStep{manager: tourManager})

	steps = append(steps, ClientNameStep{})

	return steps
}
