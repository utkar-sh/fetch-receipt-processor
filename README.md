# Receipt Processing API

The Receipt Processing API is a Go-based web service that allows you to process receipts and calculate points based on predefined rules. It provides endpoints for submitting a receipt and retrieving the points earned for a receipt.

## Installation

1. Make sure you have Go installed on your system. You can download and install it from the official Go website: https://golang.org/dl/

2. Clone the repository or download the source code files.

3. Open a terminal or command prompt and navigate to the project directory.

4. Run the following command to install the required dependencies:
   ```shell
   go mod download

1. Start the API server by running the following command:
   ```shell
   go run main.go

The server will start listening on port 8080 by default. You can change the port by modifying the code in main.go.

## Usage

### Submitting a Receipt

To submit a receipt, send a POST request to the /receipts/process endpoint with the receipt data in the request body. The receipt data should be in JSON format and include the following fields:

retailer: The name of the retailer.
* purchaseDate: The purchase date in the format yyyy-MM-dd.
* purchaseTime: The purchase time in the format HH:mm.
* total: The total amount of the purchase as a decimal number.
* items: An array of objects representing the items on the receipt. Each item object should have the following fields:
   * shortDescription: A short description of the item.
   * price: The price of the item as a decimal number.

Example request body:
   ```json
   {
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": 6.49
    },{
      "shortDescription": "Emils Cheese Pizza",
      "price": 12.25
    },{
      "shortDescription": "Knorr Creamy Chicken",
      "price": 1.26
    },{
      "shortDescription": "Doritos Nacho Cheese",
      "price": 3.35
    },{
      "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
      "price": 12.00
    }
  ],
  "total": 35.35
}
```

### Retrieving Points for a Receipt

To retrieve the points earned for a receipt, send a GET request to the /receipt/{id}/points endpoint, where {id} is the ID generated when submitting the receipt. The ID is returned in the response when a receipt is processed.

Example GET request: /receipt/abc123/points

The response will be a JSON object with the points earned for the receipt:
```json
{
  "points": 28
}
```

Please note that the provided JSON examples are for demonstration purposes, and you can adjust the request body and response format as per your application's requirements.

Feel free to explore and modify the code to suit your specific use case. If you encounter any issues or have further questions, please don't hesitate to reach out for assistance.
