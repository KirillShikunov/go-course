package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ch := make(chan int)
	ch2 := make(chan float64)
	done := make(chan struct{})

	go randomNumber(ch, 5, 10)
	go avgNumber(ch, ch2)
	go printNumber(ch2, done)

	<-done
}

func randomNumber(ch chan int, count int, max int) {
	for i := 0; i < count; i++ {
		ch <- rand.Intn(max)
	}
	close(ch)
}

func avgNumber(ch chan int, ch2 chan float64) {
	var count int
	var sum int

	for number := range ch {
		count++
		sum += number
		ch2 <- float64(sum) / float64(count)
	}
	close(ch2)
}

func printNumber(ch2 chan float64, done chan struct{}) {
	for number := range ch2 {
		fmt.Println(number)
	}
	done <- struct{}{}
}
