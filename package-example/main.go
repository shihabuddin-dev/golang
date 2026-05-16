package main // executable package

import (
	"learngo/payment"
	"learngo/test"
)

func main() {

	// bkash := Bkash{apiKey: "sdfsfd"}
	// paymentService := NewPaymentService(bkash)
	// paymentService.checkout()

	// nagad := NewNagad("safsfsf")

	mockPm := test.MockPaymentMethod{}

	paymentService1 := payment.NewPaymentService(mockPm)
	paymentService1.Checkout()

}

// module = bunch of packages
