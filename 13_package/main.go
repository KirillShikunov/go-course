package main

import "fmt"
import "github.com/KirillShikunov/compliments"

func main() {
	fmt.Println("Комплімент для твоєї дівчини:", compliments.ComplimentForWomen())
	fmt.Println("Комплімент для твого чоловіка:", compliments.ComplimentForMen())
}
