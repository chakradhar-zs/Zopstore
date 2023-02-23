
Implement a RESTful API for managing product data in memory using Go. The API should allow clients to **create**, **retrieve** and **update** products using HTTP requests to the `/products` endpoint. 
The product data should be stored in a slice of `Product` structs in memory.

### Requirements

The API should support the following HTTP request methods and operations:

#### `GET /products`

-   Retrieves a list of all products.
-   Response: `200 OK` with a JSON array of product objects in the response body.
```
[
 {
    "id": 1,
    "name": "Product A",
    "description": "A description of Product A",
    "price": 10.0
 },
 {
    "id": 2,
    "name": "Product B",
    "description": "B description of Product B",
    "price": 11.0
 }
]
```

#### `POST /products`

-   Creates a new product using the provided JSON object in the request body.
-   The new product is added to the slice of products
-   Request Body: JSON object with the following fields:
```go
{
    "id": 1,
    "name": "Product A",
    "description": "A description of Product A",
    "price": 10.0
}
```

-   Response: `201 Created` with the created product object in the response body.

### Constraints

-   All product data should be stored in a slice of `Product` structs in memory.
-   The ID field of a product must be unique.
-   Product entity : Id int ,Name string ,Description string ,Price float64