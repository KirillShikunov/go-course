package dto

import (
	"06_interface/client"
	"06_interface/tour"
	"fmt"
)

type ProcessorDTO struct {
	budget int
	client *client.Client
	tour   *tour.Tour
}

func (p *ProcessorDTO) GetBudget() int {
	return p.budget
}

func (p *ProcessorDTO) SetBudget(budget int) {
	p.budget = budget
}

func (p *ProcessorDTO) SetTour(tour *tour.Tour) {
	p.tour = tour
}

func (p *ProcessorDTO) SetClient(client *client.Client) {
	p.client = client
}

func (p *ProcessorDTO) Print() {
	fmt.Println("Інформація:")

	fmt.Printf("Тур: %s.\n", p.tour.GetName())
	fmt.Printf("Ціна: %dгрн\n", p.tour.GetPrice())
	fmt.Println("Транспорт:")
	for i, transport := range p.tour.GetTransports() {
		fmt.Printf("#%d %s\n", i+1, transport.GetName())
	}
	fmt.Printf("Ваше Ім'я: %s\n", p.client.GetName())
}
