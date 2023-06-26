package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func setupTest() {
	receipts = make(map[string]*Receipt)
}

func TestHandleProcessReceipt(t *testing.T) {
	setupTest()

	// Create a request body JSON string
	body := `{
		"retailer": "Walgreens",
		"purchaseDate": "2022-01-02",
		"purchaseTime": "08:13",
		"total": 2.65,
		"items": [
			{"shortDescription": "Pepsi - 12-oz", "price": 1.25},
			{"shortDescription": "Dasani", "price": 1.40}
		]
	}`

	// Create a new POST request with the request body
	req, err := http.NewRequest("POST", "/receipts/process", strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the handler function directly
	handleProcessReceipt(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %v, but got %v", http.StatusOK, rr.Code)
	}

	// Check the response body
	var response struct {
		ID string `json:"id"`
	}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}
	if response.ID == "" {
		t.Error("Expected non-empty response ID, but got empty")
	}
}

// Add similar test functions for other handlers

func TestMain(m *testing.M) {
	// Run setup code before running the tests
	setupTest()

	// Run the tests
	m.Run()
}
