package main

import "fmt"

func main() {
	// var numbers = [6]int{10, 20, 30} // partial array initialization
	// var orders = [6]int{10, 20, 30, 40, 50, 60}

	// slice1 := orders[1:4] // [start:end(excluded)] len, cap, pointer
	// fmt.Println(slice1)

	// // slice1[0] = 100

	// slice1 = append(slice1, 500)
	// slice1 = append(slice1, 600)
	// slice1 = append(slice1, 700)
	// slice1 = append(slice1, 800)
	// slice1 = append(slice1, 500)
	// slice1 = append(slice1, 600)
	// slice1 = append(slice1, 700)
	// slice1 = append(slice1, 800)

	// fmt.Println(slice1)
	// fmt.Println("orders", orders)

	// fmt.Println("the length of the slice is ", len(slice1))
	// fmt.Println("the cap of the slice is ", cap(slice1))

	var slice1 = []int{1, 2}
	slice1 = append(slice1, 100)
	fmt.Println(slice1)
	fmt.Println(cap(slice1))

}
