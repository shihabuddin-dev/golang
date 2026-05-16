package payment

import "fmt"

type PaymentMethod interface {
	Pay(amount float64)
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

type paymentService struct {
	method PaymentMethod
}

func NewPaymentService(method PaymentMethod) *paymentService {
	return &paymentService{
		method: method,
	}
}

func NewNagad(apiKey string) *Nagad {
	return &Nagad{
		apiKey: apiKey,
	}
}

func (ps paymentService) Checkout() {
	ps.method.Pay(10.0)
}
