package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Receipt represents a receipt data structure
type Receipt struct {
	Retailer     string  `json:"retailer"`
	PurchaseDate string  `json:"purchaseDate"`
	PurchaseTime string  `json:"purchaseTime"`
	Total        float64 `json:"total"`
	Items        []Item  `json:"items"`
}

// Item represents an item in the receipt
type Item struct {
	ShortDescription string  `json:"shortDescription"`
	Price            float64 `json:"price"`
}

// ReceiptPoints represents the points earned for a receipt
type ReceiptPoints struct {
	Points int `json:"points"`
}

// receipts stores the receipts in memory
var receipts map[string]*Receipt

func main() {
	// Initialize the receipts map
	receipts = make(map[string]*Receipt)

	// Create a new router
	router := mux.NewRouter()

	// Register the POST request handler
	router.HandleFunc("/receipts/process", handleProcessReceipt).Methods(http.MethodPost)

	// Register the GET request handler
	router.HandleFunc("/receipts/{id}/points", handleGetPoints).Methods(http.MethodGet)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", router))
}

// handleProcessReceipt handles the POST request to process a receipt
func handleProcessReceipt(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	// Parse the request body into a Receipt struct
	var receipt Receipt
	err = json.Unmarshal(body, &receipt)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Generate a new ID for the receipt
	id := uuid.New().String()

	// Save the receipt in the receipts map
	receipts[id] = &receipt

	// Create a response JSON with the generated ID
	response := struct {
		ID string `json:"id"`
	}{
		ID: id,
	}

	// Set the response status code and content type
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// Send the response JSON
	json.NewEncoder(w).Encode(response)
}

// handleGetPoints handles the GET request to retrieve the points for a receipt
func handleGetPoints(w http.ResponseWriter, r *http.Request) {
	// Get the receipt ID from the request URL
	params := mux.Vars(r)
	id := params["id"]

	// Check if the receipt ID exists in the receipts map
	receipt, exists := receipts[id]
	if !exists {
		http.NotFound(w, r)
		return
	}

	// Calculate the points for the receipt
	points := calculatePoints(receipt)

	// Create a response JSON with the points
	response := ReceiptPoints{
		Points: points,
	}

	// Set the response status code and content type
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// Send the response JSON
	json.NewEncoder(w).Encode(response)
}

// calculatePoints calculates the points earned for a receipt based on defined rules
func calculatePoints(receipt *Receipt) int {
	points := 0

	// Rule 1: One point for every alphanumeric character in the retailer name
	points += len(strings.ReplaceAll(receipt.Retailer, " ", ""))

	// Rule 2: 50 points if the total is a round dollar amount with no cents
	if math.Mod(receipt.Total, 1) == 0 {
		points += 50
	}

	// Rule 3: 25 points if the total is a multiple of 0.25
	if math.Mod(receipt.Total, 0.25) == 0 {
		points += 25
	}

	// Rule 4: 5 points for every two items on the receipt
	points += len(receipt.Items) / 2 * 5

	// Rule 5: If the trimmed length of the item description is a multiple of 3,
	// multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
	for _, item := range receipt.Items {
		trimmedLength := len(strings.TrimSpace(item.ShortDescription))
		if trimmedLength%3 == 0 {
			points += int(math.Ceil(item.Price * 0.2))
		}
	}

	// Rule 6: 6 points if the day in the purchase date is odd
	purchaseDate, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err == nil && purchaseDate.Day()%2 == 1 {
		points += 6
	}

	// Rule 7: 10 points if the time of purchase is after 2:00pm and before 4:00pm
	purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err == nil && purchaseTime.After(time.Date(0, 0, 0, 14, 0, 0, 0, time.UTC)) && purchaseTime.Before(time.Date(0, 0, 0, 16, 0, 0, 0, time.UTC)) {
		points += 10
	}

	return points
}
