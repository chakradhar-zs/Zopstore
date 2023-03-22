package middleware

import (
	"Day-19/internal/constants"
	"context"
	"net/http"
	"strings"
)

func Middle(handler http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		head := r.Header.Get("X-API-KEY")
		path := strings.Split(r.RequestURI, "/")[1]

		switch head {
		case "product-r":
			if r.Method != http.MethodGet {
				http.Error(w, "Forbidden Request", http.StatusForbidden)
				return
			}
		case "product-w":
			if r.Method != http.MethodPost {
				http.Error(w, "Forbidden Request", http.StatusForbidden)
				return
			}
		case "brand-r":
			if r.Method != http.MethodGet || strings.Contains(path, "product") {
				http.Error(w, "Forbidden Request", http.StatusForbidden)
				return
			}
		case "brand-w":
			if r.Method != http.MethodPost || strings.Contains(path, "product") {
				http.Error(w, "Forbidden Request", http.StatusForbidden)
				return
			}
		default:
			http.Error(w, "Unauthorized Request", http.StatusUnauthorized)
			return
		}
		handler.ServeHTTP(w, r)
	})
}

func MiddleOrg(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		org := r.Header.Get("oraganization")
		if r.Method == http.MethodPost {
			ctx := context.WithValue(r.Context(), constants.CtxValue, org)
			r.WithContext(ctx)
		}
		handler.ServeHTTP(w, r)
	})
}
