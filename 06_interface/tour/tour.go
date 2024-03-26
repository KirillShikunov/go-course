package tour

import (
	"fmt"
)

type Transport interface {
	Name() string
}

type Tour struct {
	id         int
	name       string
	price      int
	transports []Transport
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

func (t *Tour) Transports() []Transport {
	return t.transports
}

func NewTour(id int, name string, price int, transports []Transport) Tour {
	return Tour{id, name, price, transports}
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

func NewTourManager(tours []Tour) Manager {
	return Manager{tours}
}
