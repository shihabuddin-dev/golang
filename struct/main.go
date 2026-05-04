package main

import "fmt"

// type additionalInfo struct {
// 	phone   int
// 	address string
// }

// type user struct {
// 	name     string
// 	email    string
// 	metaInfo additionalInfo
// }

type user struct {
	name string
	age  int
	role string
}

func main() {
	// john := user{name: "john", email: "john@mail.com"}

	// john.email = "mizan@mail.com"

	// fmt.Println(john.email)

	// var user1 user

	// user1.name = "john"
	// user1.email = "test@mail.com"

	// fmt.Println(user1)

	// john := user{
	// 	name:  "john",
	// 	email: "john@mail.com",
	// 	metaInfo: additionalInfo{
	// 		phone:   017273434,
	// 		address: "165/1, Rampura",
	// 	},
	// }

	// fmt.Printf("%+v", john)

	newUser := func(name string, age int, role string) user {
		if age <= 0 {
			age = 18
		}

		return user{
			name: name,
			age:  age,
			role: role,
		}
	}

	jamal := newUser("Jamal", -25, "student")
	fmt.Println(jamal)
}

// crating struct type
// Different ways to create struct
// using printf to print struct (%+v)
// updating struct fields
// Accessing struct fields

// embedding struct
// constructor function
