package main

import (
	"net/http"
)

// Write a Go program that listens for HTTP requests on port 8080.
// The program should handle POST requests to the /products endpoint by creating a new product with the details
// provided in the request body. The request body should be a JSON object with the following fields:

// id: product id (int)
// name: the name of the product (string)
// description: a description of the product (string)
// price: the price of the product (float64)

// When a valid request is received, the program should create a new Product struct
// with the provided details and return the details of the product in the response body.
// The response should be a JSON object with the same fields as the request body.

// If the request body is missing any of the required fields or if the fields are not in the correct format,
// the program should return a 400 Bad Request status code.

// If the request method is POST set the status code to 405.

// name the entity as Product and handler as ProductHandler

func ProductHandler(w http.ResponseWriter, r *http.Request) {

}
