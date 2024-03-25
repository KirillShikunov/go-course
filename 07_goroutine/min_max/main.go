package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const MaxRandomNumber = 100

const RandomArrayLength = 10

type MinMax struct {
	min int
	max int
}

func main() {
	ch := make(chan []int)
	ch2 := make(chan MinMax)

	go goroutineFindMinMax(ch, ch2)

	for {
		go mainGoroutine(ch, ch2)
		time.Sleep(1 * time.Second)
	}
}

func mainGoroutine(ch chan []int, ch2 chan MinMax) {
	numbers := randomNumbers(MaxRandomNumber, RandomArrayLength)
	fmt.Printf("array: %v\n", numbers)
	ch <- numbers

	minMax := <-ch2
	fmt.Printf("Min: %d; Max: %d\n", minMax.min, minMax.max)
}

func goroutineFindMinMax(ch chan []int, ch2 chan MinMax) {
	for numbers := range ch {
		sort.Ints(numbers)
		ch2 <- MinMax{numbers[0], numbers[len(numbers)-1]}
	}
}

func randomNumbers(max int, length int) []int {
	numbers := make([]int, length)

	for i := 0; i < length; i++ {
		numbers[i] = rand.Intn(max)
	}

	return numbers
}
