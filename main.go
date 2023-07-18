package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)

type PaymentResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

var carts map[string]Cart

func init() {
	// Initialize the carts map
	carts = make(map[string]Cart)
}

func main() {
	http.HandleFunc("/carts/:id", handleUpdateCart)
	http.HandleFunc("/process-payment", handleProcessPayment)
	http.HandleFunc("/create-order", handleCreateOrder)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func respondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// function for handlling cart updation
func handleUpdateCart(w http.ResponseWriter, r *http.Request) {
	// Extract the cart ID from the URL parameter
	// `carts` map is used to store cart objects, where the key is the cart ID and the value is the corresponding cart object.
	cartID := strings.TrimPrefix(r.URL.Path, "/carts/")

	// Retrieve the cart from the database (assuming a map acting as a database)
	cart, found := carts[cartID]
	if !found {
		respondWithJSON(w, http.StatusNotFound, PaymentResponse{"error", "Cart not found"})
		return
	}

	// extract  the request data from the  request body to get the update cart items
	var updateReq UpdateCartRequest
	err := json.NewDecoder(r.Body).Decode(&updateReq)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, PaymentResponse{"error", "Invalid request payload"})
		return
	}

	// Updating the cart items based on the request data
	cart.Items = updateReq.Items

	// Now Save the updated cart back to the "database" (update the value in the map)
	carts[cartID] = cart

	// Respond with a success message or the updated cart
	respondWithJSON(w, http.StatusOK, PaymentResponse{"success", "Cart Updated"})
}

// funcion for payment processing
func handleProcessPayment(w http.ResponseWriter, r *http.Request) {
	// store the payment details extracted from the request body.
	var paymentReq ProcessPaymentRequest

	//decoding payment data coming as a json data
	err := json.NewDecoder(r.Body).Decode(&paymentReq)

	//error handling
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, PaymentResponse{"error", "Invalid request payload data"})
		return
	}

	//processing the payment using the payment gateway integration
	paymentGateway := NewPaymentGateway()

	isPaymentSuccessful, err := paymentGateway.ProcessPayment(paymentReq.TotalAmount, paymentReq.CardNumber, paymentReq.ExpirationDate, paymentReq.CVV)

	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, PaymentResponse{"error", "Payment Processing failed"})
		return
	}

	if !isPaymentSuccessful {
		respondWithJSON(w, http.StatusBadRequest, PaymentResponse{"error", "Payment Failed"})
		return
	}

	// Respond with a success message
	respondWithJSON(w, http.StatusOK, PaymentResponse{"success", "Payment Processed"})
}

// function to create new order
func handleCreateOrder(w http.ResponseWriter, r *http.Request) {
	//TODO
	//extract the reqeust data from the req.body to get the create order details
	var createOrderReq CreateOrderRequest

	err := json.NewDecoder(r.Body).Decode(&createOrderReq)

	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, PaymentResponse{"error", "Invalid request data payload"})
		return
	}

	// Retrieve the cart from the database (assuming a map acting as a database)
	cart, found := carts[createOrderReq.CartID]
	if !found {
		respondWithJSON(w, http.StatusNotFound, PaymentResponse{"error", "Cart not found"})
		return
	}

	// Create an instance of the Order struct
	orderData := Order{
		ID:          generateUniqueOrderID(),
		Cart:        cart,
		TotalAmount: calculateTotalAmount(cart),
		CreatedAt:   time.Now(),
	}

	// Respond with the order details
	order := CreateOrderResponse{
		OrderID:     orderData.ID,
		TotalAmount: orderData.TotalAmount,
		CreatedAt:   orderData.CreatedAt,
	}

	// Respond with the order details
	respondWithJSON(w, http.StatusOK, order)
}
