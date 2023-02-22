package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductHandler(t *testing.T) {
	successCase := Product{Id: 1, Name: "Product 1", Description: "This is product 1", Price: 10.99}
	invalidBody := map[string]interface{}{"id": "1"}

	tests := []struct {
		desc          string
		method        string
		reqBody       interface{}
		expStatuscode int
		expResp       Product
	}{
		{"success case", http.MethodPost, successCase, http.StatusCreated, successCase},
		{"error case: 405", http.MethodDelete, successCase, http.StatusMethodNotAllowed, Product{}},
		{"error case: 400", http.MethodPost, invalidBody, http.StatusBadRequest, Product{}},
	}

	for i, tc := range tests {
		body, _ := json.Marshal(tc.reqBody)

		r, err := http.NewRequest(tc.method, "/products", bytes.NewBuffer(body))
		if err != nil {
			t.Errorf("could not create request: %v", err)
			continue
		}

		w := httptest.NewRecorder()

		ProductHandler(w, r)

		assert.Equalf(t, tc.expStatuscode, w.Code, "status code mismatch")

		var p Product

		_ = json.Unmarshal(w.Body.Bytes(), &p)

		assert.Equalf(t, tc.expResp, p, "Test[%d]. Handler returned unexpected body: got %v want %v",
			i, p, tc.expResp)
	}
}
