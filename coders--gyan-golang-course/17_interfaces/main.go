package main

import "fmt"

type paymenter interface {
	pay(amount float32)
	refund(amount float32, transactionID string)
}

type payment struct {
	gateway paymenter
}

func (p *payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

func (p *payment) makeRefund(amount float32, transactionID string) {
	p.gateway.refund(amount, transactionID)
}

// razor pay
type razorpay struct {
}

func (r razorpay) pay(amount float32) {
	fmt.Println("making razorpay payment", amount)
}

func (r razorpay) refund(amount float32, transactionID string) {
	fmt.Println("making razorpay refund", amount, transactionID)
}

// stripe
type stripe struct {
}

func (r stripe) pay(amount float32) {
	fmt.Println("making stripe payment", amount)
}

func (r stripe) refund(amount float32, transactionID string) {
	fmt.Println("making stripe refund", amount, transactionID)
}

func main() {

	// stripeGateway := stripe{}
	razorpayGateway := razorpay{}

	newPayment := payment{
		gateway: razorpayGateway,
	}

	newPayment.makePayment(100)
	newPayment.makeRefund(50, "txn_123")
}
