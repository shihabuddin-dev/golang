package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var fileUrl string

func main() {
	var start = time.Now()

	// wg.Add(1) // 1
	// go uploadFile()

	// wg.Add(1) // 2
	// go saveToDB()
	// wg.Add(1) // 3
	// go sendEmail()

	// uploadFile()
	// saveToDB()
	// sendEmail()

	wg.Go(uploadFile)
	wg.Go(saveToDB)
	wg.Go(sendEmail)

	wg.Wait() // waiting .... until counter is 0

	fmt.Println("file url", fileUrl)
	fmt.Println("all tasks completed")
	fmt.Println("time taken", time.Since(start))

}

func uploadFile() {
	fmt.Println("uploading file")
	time.Sleep(3 * time.Second)
	fmt.Println("file uploaded successfully")

}

func saveToDB() {
	fmt.Println("saving to db")
	time.Sleep(1 * time.Second)
	fmt.Println("saved to db successfully")

}

func sendEmail() {
	fmt.Println("sending email")
	time.Sleep(2 * time.Second)
	fmt.Println("email sent successfully")

}
