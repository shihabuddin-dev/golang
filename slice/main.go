package main

import "fmt"

func main() {
	var numbers = [5]int{2, 3, 4, 8, 2}

	sliceNumbers := numbers[1:4]
	sliceNumbers = append(sliceNumbers, 30)

	fmt.Println(sliceNumbers)
	fmt.Println("Total Length of Numbers", len(sliceNumbers))
	fmt.Println("Total Capacity of Numbers", cap(sliceNumbers))
}
