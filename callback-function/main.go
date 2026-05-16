package main

import "fmt"

// callback function
// first class citizen =>

// func process(sayHello func()) {
// 	sayHello()
// }

// callback function is a function that takes a function as a parameter
//  যে ফাংশনটিকে অন্য কোনো ফাংশনের ভেতর আর্গুমেন্ট হিসেবে পাস করা হয়, তাকে কলব্যাক ফাংশন বলা হয়
func calculate(a int, b int, operation func(x int, y int) int) int {
	return operation(a, b)
}

func main() {

	// greet := func() {
	// 	fmt.Println("Hello there")
	// }
	// process(greet)

	add := func(n1 int, n2 int) int {
		return n1 + n2
	}
	multiply := func(n1 int, n2 int) int {
		return n1 * n2
	}
	fmt.Println(calculate(10, 20, add))
	fmt.Println(calculate(10, 20, multiply))

	// anonymous callback function
	result := calculate(10, 5, func(x int, y int) int {
		return x - y
	})

	fmt.Println(result)
}
