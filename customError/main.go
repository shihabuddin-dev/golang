package main

import "fmt"

type CustomError struct {
	message string
	code    int
}

func (cu *CustomError) Error() string {
	return cu.message
}

func login(password string) error {
	if password != "1234" {
		return &CustomError{
			message: "Password do not match",
			code:    401,
		}
	}
	return nil
}

func main() {
	err := login("2250")
	if err != nil {

		customErr, ok := err.(*CustomError)

		if !ok {
			fmt.Println(customErr)
		} else {

			fmt.Println(customErr)
			fmt.Println(customErr.code)
		}

		// if customError, ok := err.(*CustomError); ok {
		// 	fmt.Println("Getting Error With", customError.code)
		// }

		// fmt.Println("Error", err, "Code", err.code)
	}

	users := map[int]string{
		1: "Hanki",
		2: "Panki",
		3: "Danki",
	}
	name, ok := users[2] // ""

	fmt.Println(name)
	fmt.Println(ok)

	if ok {
		// do something
		fmt.Println("User found")
	}

	fmt.Println("Main Codes")
}
