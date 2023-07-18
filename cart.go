package main

type CartItem struct {
	ID       string  `json:"id"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type Cart struct {
	ID    string     `json:"id"`
	Items []CartItem `json:"items"`
}
