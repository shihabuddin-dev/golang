package main

import (
	"fmt"
	"log"
)

func doSomething() {

	defer func() {
		fmt.Println("Defer function ran")
		r := recover()
		if r != nil {
			fmt.Println("recovered form", r)
		}
	}()
	panic("something really bad happen") // missile (recoverable)

}

func doAnotherThing() {
	defer func() {
		fmt.Println("Defer function ran")
	}()

	log.Fatal("something very big happen") // atomic boom (terminate)
}

func main() {
	// doSomething()

	doAnotherThing()

	fmt.Println("Main Completed Normally")
}
