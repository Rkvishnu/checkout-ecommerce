package main

type UpdateCartRequest struct {
	Items []CartItem `json:"items"`
}

type ProcessPaymentRequest struct {
	CartID         string  `json:"cartID"`
	CardNumber     string  `json:"cardNumber"`
	ExpirationDate string  `json:"expirationDate"`
	CVV            string  `json:"cvv"`
	TotalAmount    float64 `json:"totalAmount"`
}

type CreateOrderRequest struct {
	CartID string `json:"cartID"`
}
