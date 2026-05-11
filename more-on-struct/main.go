package main

import "fmt"

// type user struct {
// 	name       string
// 	age        int
// 	isLoggedIn bool
// 	greet      func()
// }

type user struct {
	name       string
	age        int
	isLoggedIn bool
}

func main() {
	// person := user{
	// 	name:       "shihab",
	// 	age:        22,
	// 	isLoggedIn: true,
	// }

	// person.greet = func() {
	// 	fmt.Println("Hello ", person.name)
	// }
	// fmt.Println(person)
	// person.greet()

	user1 := user{
		name:       "Boss",
		age:        30,
		isLoggedIn: false,
	}
	user2 := user{
		name:       "Original",
		age:        30,
		isLoggedIn: false,
	}
	user1.greet()
	user2.greet()
	fmt.Printf("%+v", user1)
}

// receiver function
func (u user) greet() {
	fmt.Println("Hello!", u.name)
}

func (u *user) login() {
	fmt.Println("Login called")
	u.isLoggedIn = true
}

func (u *user) logOut() {
	fmt.Println("Log Out Success", u.name)
	u.isLoggedIn = false
}
