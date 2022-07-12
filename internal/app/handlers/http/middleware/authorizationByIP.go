package middleware

import (
	"net"
	"net/http"
)

func AuthorizationByIP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, IPNet, err := net.ParseCIDR("127.0.0.1")
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		if !IPNet.Contains(net.ParseIP(r.Header.Get("X-Real-IP"))) {
			http.Error(w, "Your subnet is forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
