package main

import "fmt"

// Implement a function in Go that retrieves the contents from a URL using the HTTP client package, and
// returns the response body as a string.

// Hints:

/*
- Use the http.Get() function from the net/http package to retrieve the response from the URL.
- Check the response status code before reading the response body.
- Use a buffer to read the response body.
- Ensure to close the response body to prevent resource leaks.
*/

func getRespBody(url string) string {
	return ""
}

func main() {
	body := getRespBody("https://jsonplaceholder.typicode.com/todos/1")

	expected := `{
  "userId": 1,
  "id": 1,
  "title": "delectus aut autem",
  "completed": false
}`

	if body != expected {
		fmt.Println("expected:%s,\n got:%s", expected, body)
	}

	fmt.Println("successful")
}
