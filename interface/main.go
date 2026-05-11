package main

import (
	"fmt"
)

// ts => interface for data shape
// go => interface behavior contract

type Animal interface {
	speak()
}

type Dog struct{}
type Cat struct{}
type Human struct {
	name string
}

// dog er
func (d Dog) speak() {
	fmt.Println("Woof!!!")
}

// cat er
func (c Cat) speak() {
	fmt.Println("Me ao!!!")
}

func (h Human) speak() {
	fmt.Println("My name is", h.name)
}

func makeSound(a Animal) {
	a.speak()
}

// abstraction, polymorphism, inheritance and encapsulation

func main() {

	dexter := Dog{}
	billi := Cat{}
	boss := Human{name: "boss"}

	boss.name = "Muslim Cat"

	makeSound(dexter)
	makeSound(billi)
	makeSound(boss)
}
