package middleware

import (
	"github.com/PanovAlexey/url_carver/config"
	"net"
	"net/http"
)

func AuthorizationByIP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		trustedSubNet := (config.New()).GetTrustedSubNet()

		_, IPNet, err := net.ParseCIDR(trustedSubNet)

		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		if len(trustedSubNet) == 0 || !IPNet.Contains(net.ParseIP(r.Header.Get("X-Real-IP"))) {
			http.Error(w, "Your subnet is forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
