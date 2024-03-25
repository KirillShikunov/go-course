package tour

import (
	"06_interface/transport"
	"fmt"
)

type Tour struct {
	id         int
	name       string
	price      int
	transports []transport.Transport
}

func (t *Tour) Price() int {
	return t.price
}

func (t *Tour) Id() int {
	return t.id
}

func (t *Tour) Name() string {
	return t.name
}

func (t *Tour) Transports() []transport.Transport {
	return t.transports
}

type Manager struct {
	tours []Tour
}

func (m *Manager) FindByBudget(budget int) []Tour {
	var budgetTours []Tour

	for _, tour := range m.tours {
		if tour.price <= budget {
			budgetTours = append(budgetTours, tour)
		}
	}

	return budgetTours
}

func (m *Manager) Get(id int) (*Tour, error) {
	for _, tour := range m.tours {
		if tour.id == id {
			return &tour, nil
		}
	}

	return nil, fmt.Errorf("тур(#%d) не вдалося знайти", id)
}

func BuildTourManager() Manager {
	var tours []Tour

	car := transport.NewCar("456DFG")
	plane := transport.NewPlane("Airbus A320")
	boat := transport.NewBoat("Лотка з ухилятнами")
	train := transport.NewTrain("Eurostar")

	tours = append(tours, Tour{1, "Вроцлав - Відень", 500, []transport.Transport{car, plane}})
	tours = append(tours, Tour{2, "Прага - Берлін", 400, []transport.Transport{car, train}})
	tours = append(tours, Tour{3, "Київ - Угорщина(через Тису)", 3000, []transport.Transport{car, boat}})

	return Manager{tours}
}
