package main

import "fmt"

// Variadic function

func add(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total = total + number
	}

	return total

}

func greet(prefix string, mps ...string) {
	for _, mp := range mps {
		fmt.Println(prefix, mp)
	}
}

func main() {

	sum := add(5, 10, 20, 50)
	fmt.Println(sum)

	mps := []string{"Jamal", "Kamal", "Khairul", "Milon", "Hero Alam"}

	greet("Welcome", mps...) // []string{"Jamal", "Kamal", "Khairul"} => "Jamal", "Kamal", "Khairul"

}

// flexible amount of argument
// must be the last parameter
// internally slice
