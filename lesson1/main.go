package main

import "fmt"

func main() {
	fmt.Print("Всім привіт!")
	fmt.Printf("Мене звати, %s.\n", "Кирило")
	fmt.Printf("Мені %d років.\n", 25)
	fmt.Printf("У мене є два брати: %v.\n", [2]string{"Миролав", "Артем"})
	fmt.Printf("Я наймолодший серед них. І це %t", true)
}
