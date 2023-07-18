package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Order struct {
	ID          string    `json:"id"`
	Cart        Cart      `json:"cart"`
	TotalAmount float64   `json:"totalAmount"`
	CreatedAt   time.Time `json:"createdAt"`
}

func generateUniqueOrderID() string {
	// Example: Use a combination of timestamp, random number, and user ID
	return fmt.Sprintf("%d-%d", time.Now().Unix(), rand.Intn(10000))
}

// Calculating the total amount for a given cart
func calculateTotalAmount(cart Cart) float64 {
	total := 0.0
	for _, item := range cart.Items {
		total += item.Price * float64(item.Quantity)
	}
	return total
}