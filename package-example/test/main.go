package test

import "fmt"

type MockPaymentMethod struct {
}

func (mockPM MockPaymentMethod) Pay(amount float64) {
	// test logic
	fmt.Println("testing payment method")
}
