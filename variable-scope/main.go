package main

import "fmt"

func main() {
	sugar := 2

	// annoymous function
	makeCoffee := func() {
		fmt.Println("Coffee is making")
	}
	makeCoffee()

	// iife function
	func(coffeeType string) {
		fmt.Printf("%s Coffee is making \n", coffeeType)
	}("Black")

	makingCoffee := func() {
		sugar := 30
		coffee := "Cappuccino"
		fmt.Printf("%s Coffee is making with sugar %d \n", coffee, sugar)
	}
	makingCoffee()

	fmt.Printf("my sugar is %d gram", sugar)
}
