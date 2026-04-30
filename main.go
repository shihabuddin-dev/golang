package main

import "fmt"

func makeCoffee(kind string, isSugar bool) {
	fmt.Println("Making Coffee......")
	fmt.Printf("%s Coffee is ready with Sugar %t \n", kind, isSugar)
}

func makeTea(kind string) (tea string, price int) {
	price = 10
	tea = fmt.Sprintf("%s Tea is ready", kind)
	return
}

func main() {

	// **variable declaration**
	// name := "shihab-uddin" // short way, mostly used
	// var name = "shihab-uddin" // long way without type
	// var name string = "Shihab Uddin" // long way with type

	// grouped variable
	var (
		name      string = "shihab-uddin"
		age       int    = 23
		isMarried bool   = true
	)

	// multiple variables
	// var x, y int
	// x = 10
	// y = 20

	var x, y int = 10, 20

	var a, b string = "shihab", "uddin"

	var point int = 4

	point = 5

	const pi = 3.1416

	// default value of variable
	var score int
	var nameOfScore string
	var isRealScore bool
	var isRealScoreFloat float64

	result := 6.25

	fmt.Println(name, age, isMarried)
	fmt.Println(x, y)
	fmt.Println(a, b)

	fmt.Println(point)
	fmt.Println(pi)
	fmt.Println("Default Value", score, nameOfScore, isRealScore, isRealScoreFloat)

	fmt.Printf("I am %s and I am %d years old", name, age)
	fmt.Printf("I am %s and my score is %.2f", name, result)
	formatString := fmt.Sprintf("I am %s and my score is %.2f", name, result)
	fmt.Println(formatString)

	// run function
	makeCoffee("black", true)
	makeCoffee("cold", false)

	myTea, bill := makeTea("Milk")
	// myTea2 := makeTea("Green")
	fmt.Printf("%s and bill is %d TK", myTea, bill)
	// fmt.Println(myTea2)

}
