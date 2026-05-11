package main

import "fmt"

type PaymentMethod interface {
	pay(amount float64)
}

type Bkash struct {
	apiKey string
}

type Nagad struct {
	apiKey string
}

func (bk *Bkash) pay(amount float64) {
	// actual payment logic
	fmt.Printf("Paying %.2f tk with Bkash\n", amount)
}
func (ng *Nagad) pay(amount float64) {
	// actual payment logic
	fmt.Printf("Paying %.2f tk with Nagad", amount)
}

type PaymentService struct {
	method PaymentMethod
}

func NewPaymentService(method PaymentMethod) *PaymentService {
	return &PaymentService{
		method: method,
	}
}

func NewNagad(apiKey string) *Nagad {
	return &Nagad{
		apiKey: apiKey,
	}
}

func (ps PaymentService) checkout() {
	ps.method.pay(10.0)
}

type MockPaymentMethod struct {
}

func (mockPM MockPaymentMethod) pay(amount float64) {
	// test logic
	fmt.Println("testing payment method")
}

func main() {

	// bkash := Bkash{apiKey: "sdfsfd"}
	// paymentService := NewPaymentService(bkash)
	// paymentService.checkout()

	// nagad := NewNagad("safsfsf")

	mockPm := MockPaymentMethod{}

	paymentService1 := NewPaymentService(mockPm)
	paymentService1.checkout()

}
