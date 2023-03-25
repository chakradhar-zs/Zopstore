package middleware

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"Day-19/internal/constants"
)

// TestMiddle is a test function which uses a mock handler to test Middle function
func TestMiddle(t *testing.T) {
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})

	tests := []struct {
		desc      string
		method    string
		key       string
		path      string
		expStatus int
	}{
		{"Success", http.MethodGet, "product-r", "/products", http.StatusOK},
		{"Success", http.MethodPost, "product-w", "/products", http.StatusOK},
		{"Fail", http.MethodPost, "product-r", "/products", http.StatusForbidden},
		{"Fail", http.MethodPost, "brand-w", "/products", http.StatusForbidden},
		{"Fail", http.MethodGet, "", "/products", http.StatusUnauthorized},
		{"Fail", http.MethodGet, "product-w", "/products", http.StatusForbidden},
		{"Fail", http.MethodPost, "brand-r", "/brands", http.StatusForbidden},
	}

	for i, val := range tests {
		h := Middle(nextHandler)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(val.method, val.path, http.NoBody)
		req.Header.Set("X-API-KEY", val.key)
		h.ServeHTTP(w, req)
		assert.Equalf(t, w.Code, val.expStatus, "Test[%d] Failed. \n%s", i, val.desc)
	}
}

// TestMiddleOrg is a test function which uses a mock handler to test MiddleOrg function
func TestMiddleOrg(t *testing.T) {
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})

	tests := []struct {
		desc   string
		method string
		org    string
		path   string
		expKey string
	}{
		{"Success", http.MethodPost, "zs", "/products", "zs"},
		{"Success", http.MethodPost, "", "/products", ""},
		{"Success", http.MethodGet, "zs", "/products", ""},
	}
	for i, val := range tests {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(val.method, val.path, http.NoBody)
		ctx := context.WithValue(req.Context(), constants.CtxValue, val.expKey)
		req = req.WithContext(ctx)
		req.Header.Set("X-ORG", val.org)

		h := MiddleOrg(nextHandler)
		h.ServeHTTP(w, req)

		out := req.Context().Value(constants.CtxValue)

		assert.Equalf(t, out, val.expKey, "Test[%d] Failed. \n%s", i, val.desc)
	}
}
