package middleware

import (
	"github.com/google/uuid"
	"net/http"
	"time"
)

const userTokenName = `token`
const userTokenCookieExpirationDate = 1 * time.Minute

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userToken := getUserTokenFromCookie(r)

		if !isUserTokenValid(userToken) {
			userToken = userTokenGenerate()
			setUserTokenToCookie(userToken, w)
		}

		next.ServeHTTP(w, r)
	})
}

func getUserTokenFromCookie(r *http.Request) string {
	userTokenCookie, err := r.Cookie(userTokenName)

	if err != nil {
		if err != http.ErrNoCookie {
			// @ToDo: log some error oO
		}

		return ``
	}

	return (*userTokenCookie).Value
}

func isUserTokenValid(userToken string) bool {
	if len(userToken) < 1 {
		return false
	}

	return true
}

func userTokenGenerate() string {
	return uuid.New().String()
}

func setUserTokenToCookie(userToken string, w http.ResponseWriter) {
	cookie := http.Cookie{
		Name:    userTokenName,
		Value:   userToken,
		Expires: time.Now().Add(userTokenCookieExpirationDate),
	}

	http.SetCookie(w, &cookie)
}
