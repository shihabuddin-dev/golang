package main

import "fmt"

func ageStatus(age int) {
	if age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Child")
	}
}

func main() {
	age := 20

	if age >= 18 {
		fmt.Println("Adult Person")
	} else {
		fmt.Println("Child Person")
	}

	ageStatus(15)
	ageStatus(30)
	fmt.Print("hello boss \n")

	// new style declare variable with if-else score := 85; it will work for only condition block

	if score := 85; score >= 80 {
		fmt.Println("You Got A+ and Your score is", score)
	} else if score >= 70 {
		fmt.Println("You Got A and Your score is", score)
	} else if score >= 60 {
		fmt.Println("You Got B and Your score is", score)
	} else {
		fmt.Println("Failed to Exam and Your score is", score)
	}
}
