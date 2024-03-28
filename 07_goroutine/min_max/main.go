package main

import (
	"fmt"
	"math"
	"math/rand"
)

type MinMax struct {
	Min int
	Max int
}

func main() {
	inputCh := make(chan int)
	outputCh := make(chan MinMax)
	done := make(chan struct{})

	go randomNumber(inputCh, outputCh, done)
	go findMinMax(inputCh, outputCh)

	<-done
}

func randomNumber(inputCh chan int, outputCh chan MinMax, done chan struct{}) {
	for i := 0; i < 10; i++ {
		inputCh <- rand.Intn(100)
	}
	close(inputCh)

	minMax := <-outputCh
	fmt.Printf("Min number: %d, Max number: %d\n", minMax.Min, minMax.Max)

	close(outputCh)

	done <- struct{}{}
}

func findMinMax(inputCh chan int, outputCh chan MinMax) {
	var numbers []int
	for number := range inputCh {
		fmt.Printf("Random number: %d\n", number)
		numbers = append(numbers, number)
	}

	if len(numbers) == 0 {
		outputCh <- MinMax{}
		return
	}

	minNumber := math.MaxInt64
	maxNumber := math.MinInt64
	for _, number := range numbers {
		if number < minNumber {
			minNumber = number
		}
		if number > maxNumber {
			maxNumber = number
		}
	}

	outputCh <- MinMax{Min: minNumber, Max: maxNumber}
}
