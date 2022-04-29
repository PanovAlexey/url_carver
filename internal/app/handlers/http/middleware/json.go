package middleware

import (
	"net/http"
)

func JSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestContentType := r.Header.Get("Content-Type")

		if requestContentType != "application/json" {
			http.Error(w, "Invalid Content-Type", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}
