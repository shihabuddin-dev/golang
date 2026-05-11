package main

import "fmt"

func change(x *int) {
	*x = 5000
	fmt.Println("inside function:", *x)

}

func main() {

	// a := 42
	// // p := a // 42
	// p := &a // 0xc000012138 -> pointer

	// a = 500

	// // & -> address of a variable
	// // * -> dereference (value from address)

	// *p = 1000

	// fmt.Println("a:", a)
	// fmt.Println("p:", *p)
	// fmt.Println("p:", p)

	// y := 20

	// change(&y)

	// fmt.Println("outside function y:", y)

	bigArray := [5]int{10, 20, 30, 40, 50}

	// modifyWithoutPointer(bigArray)
	// fmt.Println("After without pointer:", bigArray)

	modifyWithPointer(&bigArray)
	fmt.Println("After with pointer:", bigArray)

}

func modifyWithoutPointer(arr [5]int) {
	arr[0] = 999
	fmt.Println("Inside without pointer:", arr)
}

func modifyWithPointer(arr *[5]int) {
	arr[0] = 777
	fmt.Println("Inside with pointer:", *arr)
}

// Pointer basics
// Dereference Example
// Modify Value Using Pointer
// Pointer in Function
// Nil pointer

// Large data copy avoid করা যায়
// Struct modify করা যায়
// Performance improve হয়
