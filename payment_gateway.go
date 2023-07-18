package main

import "fmt"

type PaymentGateway struct {
}

func NewPaymentGateway() *PaymentGateway {
	return &PaymentGateway{}
}

func (pg *PaymentGateway) ProcessPayment(amount float64, cardNumber string, expirationDate string, cvv string) (bool, error) {
	// This is a dummy implementation that considers any payment amount greater than 0 as successful
	if amount > 0 {
		return true, nil
	}

	return false, fmt.Errorf("payment failed")
}
