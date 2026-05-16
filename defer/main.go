// package main

// import "fmt"

// // func deferred(result int) {
// // 	fmt.Println("deferred result:", result)
// // }

// func example() int {
// 	result := 10

// 	defer func() {
// 		result = result + 200
// 		fmt.Println("deferred result:", result) // 110
// 	}()

// 	fmt.Println("I am from example fn:", result) // 10

// 	// result += 100
// 	result = result + 100

// 	return result // 110
// }

// func main() {
// 	fmt.Println("return result", example())
// }

// I am from example fn: 10
// deferred result: 310
// return result 110

// ====================
// name return

// I am from example fn: 10
// deferred result: 310
// return result 310

package main

import "fmt"

func example() int {
	result := 10

	defer func() {
		result = result + 200
		fmt.Println("deferred result:", result) // 310
	}()

	fmt.Println("I am from example fn:", result) // 10

	// result += 100
	result = result + 100

	return result // 110
}

func namedExample() (result int) {
	result = 10

	defer fmt.Println("Named example function exited...")
	defer func() {
		result = result + 200
		fmt.Println("deferred result:", result) // 310
	}()
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// defer fmt.Println(3)
	// defer fmt.Println(4)

	fmt.Println("I am from example fn:", result) // 10

	// result += 100
	result = result + 100

	return // 0x545345345
}

func connectDb() {
	defer fmt.Println("Closing database connection...")
	fmt.Println("Connecting to database...")
	// some operations

}

func main() {
	fmt.Println("named return result", namedExample()) // 310
	fmt.Println("===================================")
	fmt.Println("return result", example()) // 110

	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	defer fmt.Println(4)

}
