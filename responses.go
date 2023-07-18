package main

import "time"

type UpdateCartResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ProcessPaymentResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type CreateOrderResponse struct {
	OrderID     string    `json:"orderID"`
	TotalAmount float64   `json:"totalAmount"`
	CreatedAt   time.Time `json:"createdAt"`
}
