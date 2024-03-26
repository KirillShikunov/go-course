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

func NewBudgetStep() BudgetStep {
	return BudgetStep{}
}

type TourStep struct {
	manager tour.Manager
}

func NewTourStep(manager tour.Manager) TourStep {
	return TourStep{manager: manager}
}

func (t TourStep) Print(dto *dto.ProcessorDTO) error {
	fmt.Println("Тури за вашим бюджетом:")
	toursByBudget := t.manager.FindByBudget(dto.Budget())
	if len(toursByBudget) == 0 {
		return fmt.Errorf("турів за вашим бюджетом не знайдено")
	}

	for _, tourBudget := range toursByBudget {
		fmt.Printf("#%d %s | Ціна: %dгрн\n", tourBudget.Id(), tourBudget.Name(), tourBudget.Price())
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

type ClientStep struct {
}

func (n ClientStep) Print(dto *dto.ProcessorDTO) error {
	fmt.Println("Вкажіть ваше ім'я:")
	return nil
}

func (n ClientStep) Process(value string, dto *dto.ProcessorDTO) error {
	newClient := client.NewClient(value)
	dto.SetClient(&newClient)
	return nil
}

func NewClientStep() ClientStep {
	return ClientStep{}
}
