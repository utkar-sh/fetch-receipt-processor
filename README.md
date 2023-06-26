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

### Submitting a ReceiptTo submit a receipt, send a POST request to the /receipts/process endpoint with the receipt data in the request body. The receipt data should be in JSON format and include the following fields:

retailer: The name of the retailer.
purchaseDate: The purchase date in the format yyyy-MM-dd.
purchaseTime: The purchase time in the format HH:mm.
total: The total amount of the purchase as a decimal number.
items: An array of objects representing the items on the receipt. Each item object should have the following fields:
shortDescription: A short description of the item.
price: The price of the item as a decimal number.
Example request body:
