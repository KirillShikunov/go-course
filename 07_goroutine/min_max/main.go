package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	ch := make(chan int)
	done := make(chan struct{})

	go randomNumber(ch, 5, 10)
	go findMinMax(ch, done)

	<-done
}

func randomNumber(ch chan int, count int, max int) {
	for i := 0; i < count; i++ {
		ch <- rand.Intn(max)
	}
	close(ch)
}

func findMinMax(ch chan int, done chan struct{}) {
	var numbers []int
	for number := range ch {
		fmt.Printf("Random number: %d\n", number)
		numbers = append(numbers, number)
	}

	sort.Ints(numbers)
	minNumber := numbers[0]
	maxNumber := numbers[len(numbers)-1]

	fmt.Printf("Min: %d; Max: %d", minNumber, maxNumber)

	done <- struct{}{}
}
