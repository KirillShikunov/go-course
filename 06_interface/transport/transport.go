package transport

import "fmt"

type Transport interface {
	Name() string
}

type Car struct {
	number string
}

func (c *Car) Name() string {
	return fmt.Sprintf("Авто №%s", c.number)
}

func NewCar(number string) *Car {
	return &Car{number: number}
}

type Plane struct {
	name string
}

func (p *Plane) Name() string {
	return fmt.Sprintf("Літак \"%s\"", p.name)
}

func NewPlane(name string) *Plane {
	return &Plane{name: name}
}

type Boat struct {
	name string
}

func (b *Boat) Name() string {
	return fmt.Sprintf("Човен \"%s\"", b.name)
}

func NewBoat(name string) *Boat {
	return &Boat{name: name}
}

type Train struct {
	name string
}

func (t *Train) Name() string {
	return fmt.Sprintf("Поїзд \"%s\"", t.name)
}

func NewTrain(name string) *Train {
	return &Train{name: name}
}
