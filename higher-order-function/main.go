package main

import "fmt"

// higher order function is a function that takes a function as a parameter
// যে ফাংশন অন্য কোনো ফাংশনকে প্যারামিটার হিসেবে গ্রহণ করে অথবা কোনো ফাংশনকে রিটার্ন করে, তাকে হায়ার অর্ডার ফাংশন বলে।
func calculate(a int, b int, operation func(x int, y int) int) int {
	return operation(a, b)
}

func multiplyBy(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func main() {
	// add := func(n1 int, n2 int) int {
	// 	return n1 + n2
	// }
	// fmt.Println(calculate(10, 20, add))

	double := multiplyBy(2) //func(x int) int {return x * factor}
	fmt.Println(double(5))
	fmt.Println(double(555))

	triple := multiplyBy(3)
	fmt.Println(triple(344))

}
