package main

import "fmt"

type Animal struct {
	name  string
	class string
}

func (animal *Animal) fullName() string {
	return fmt.Sprintf("%s(%s)", animal.name, animal.class)
}

type Cage struct {
	animal *Animal
}

type Zookeeper struct {
	cageManager *CageManager
}

func (zookeeper *Zookeeper) dump() {
	zookeeper.cageManager.dump()
}

const MaxCagesNumber = 5

func (zookeeper *Zookeeper) catchAnimal(animal *Animal) {
	zookeeper.cageManager.addAnimal(animal)
}

func main() {
	cageManager := CageManager{make([]Cage, 0, MaxCagesNumber)}
	zookeeper := Zookeeper{&cageManager}

	elephant := Animal{"Mike", "Elephant"}
	lion := Animal{"Alex", "lion"}

	zookeeper.catchAnimal(&elephant)
	zookeeper.catchAnimal(&lion)
	zookeeper.dump()
}

type CageManager struct {
	cages []Cage
}

func (manager *CageManager) addAnimal(animal *Animal) {
	if manager.isFreeSpot() == false {
		fmt.Printf("Error: You cannot add %s to cage because do not have free cage. \n", animal.fullName())
		return
	}

	manager.cages = append(manager.cages, Cage{animal})
}

func (manager *CageManager) isFreeSpot() bool {
	return len(manager.cages) < MaxCagesNumber
}

func (manager *CageManager) dump() {
	fmt.Print("Cages status:\n")
	for i, cage := range manager.cages {
		fmt.Printf("%s in a cage #%d. \n", cage.animal.fullName(), i+1)
	}
}
