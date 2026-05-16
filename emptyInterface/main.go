package main

import "fmt"

// any = interface{}
// type assertion = interface{}
// ok idiom

func Process(data any) {
	// fmt.Print(data)

	strData, ok := data.(string)
	if ok {
		fmt.Println(len(strData))
	}

	intData, ok := data.(int)
	if ok {
		fmt.Println(intData + 100)
	}
}

func main() {

	// Process([]int{6, 7, 8, 9})
	// var data interface{}
	// data = "boss"
	// fmt.Println(data)

	// data = 20
	// fmt.Println(data)
	Process(10)

}
