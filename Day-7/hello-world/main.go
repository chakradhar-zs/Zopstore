package main

import "net/http"

// Implement unit test for the below handler

// hello returns hello-world for a GET method.
func hello(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}

	w.WriteHeader(http.StatusOK)

	w.Write([]byte("hello-world"))
}
