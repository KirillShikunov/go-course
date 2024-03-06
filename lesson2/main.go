package main

import "fmt"

type Animal struct {
	name  string
	class string
}

type Cage struct {
	animal Animal
	//number int
}

type Zookeeper struct {
	cage *Cage
}

func (zookeeper *Zookeeper) dump() {
	if zookeeper.cage == nil {
		fmt.Print("Cage is empty")
		return
	}

	fmt.Printf("%s(%s) is in a cage", zookeeper.cage.animal.class, zookeeper.cage.animal.name)
}

func (zookeeper *Zookeeper) pushAnimal(animal Animal) {
	if zookeeper.cage != nil {
		fmt.Print("Cage is not empty")
		return
	}
	zookeeper.cage = &Cage{animal}

}

func (zookeeper *Zookeeper) popAnimal(animal Animal) {
	zookeeper.cage = &Cage{animal}
}

func main() {
	elephant := Animal{"Mike", "Elephant"}
	zookeeper := Zookeeper{}

	zookeeper.pushAnimal(elephant)
	zookeeper.dump()
}
