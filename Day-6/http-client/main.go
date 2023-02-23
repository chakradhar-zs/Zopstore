package main

// Implement a function in Go that retrieves the contents from a URL using the HTTP client package, and
// returns the response body as a string.

// Hints:

/*
- Use the http.Get() function from the net/http package to retrieve the response from the URL.
- Check the response status code for errors before reading the response body.
- Use a buffer to read the response body.
- Ensure to close the response body to prevent resource leaks.
- Handle errors gracefully by returning an error value in case of any failures.
*/

func getRespBody(url string) (string, error) {
	return "", nil
}
