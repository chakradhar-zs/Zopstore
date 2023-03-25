package middleware

import (
	"context"
	"net/http"
	"strings"

	"Day-19/internal/constants"
)

func isPresent(m string, ms []string) bool {
	for _, v := range ms {
		if m == v {
			return true
		}
	}

	return false
}

// Middle takes http handler and checks for X-API-KEY
func Middle(handler http.Handler) http.Handler {
	apikeys := make(map[string][]string)
	apikeys["product-r"] = []string{"GET", "products", "brands"}
	apikeys["product-w"] = []string{"POST", "PUT", "products", "brands"}
	apikeys["brand-r"] = []string{"GET", "brands"}
	apikeys["brand-w"] = []string{"POST", "PUT", "brands"}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		head := r.Header.Get("X-API-KEY")
		key, ok := apikeys[head]

		if !ok {
			http.Error(w, "Unauthorized Request", http.StatusUnauthorized)
			return
		}

		if !isPresent(r.Method, key) {
			http.Error(w, "Forbidden Request", http.StatusForbidden)
			return
		}

		path := strings.Split(r.URL.String(), "/")[1]
		path = strings.Split(path, "?")[0]

		if !isPresent(path, key) {
			http.Error(w, "Forbidden Request", http.StatusForbidden)
			return
		}
		handler.ServeHTTP(w, r)
	})
}

// MiddleOrg takes http handler and adds X-ORG to request
func MiddleOrg(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		org := r.Header.Get("X-ORG")
		if r.Method == http.MethodPost || r.Method == http.MethodPut {
			ctx := context.WithValue(r.Context(), constants.CtxValue, org)
			r = r.WithContext(ctx)
		}
		if r.Method == http.MethodGet {
			url := r.URL
			query := url.Query()
			query.Add("organization", org)
			r.URL.RawQuery = query.Encode()
		}
		handler.ServeHTTP(w, r)
	})
}
