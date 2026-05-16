package main

import "fmt"

// closure is a function that returns another function as a value

func makeCounter() func() int {
	count := 0

	inner := func() int {
		count++
		return count
	}

	return inner
}

func main() {
	next := makeCounter()

	fmt.Println(next()) // 1
	fmt.Println(next()) // 2
	fmt.Println(next()) // 3
}
