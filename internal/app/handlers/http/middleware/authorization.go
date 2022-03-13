package middleware

import (
	"github.com/PanovAlexey/url_carver/internal/app/services"
	"github.com/google/uuid"
	"log"
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
			encryptionService, err := services.NewEncryptionService()

			if err == nil {
				userTokenEncrypted := encryptionService.Encrypt(userToken)
				setUserTokenToCookie(userTokenEncrypted, w)
			} else {
				setUserTokenToCookie(userToken, w)
			}

		}

		next.ServeHTTP(w, r)
	})
}

func getUserTokenFromCookie(r *http.Request) string {
	userToken := ``
	userTokenCookie, err := r.Cookie(userTokenName)

	if err != nil {
		if err != http.ErrNoCookie {
			log.Println("error with getting token from cookie: " + err.Error())
		}

		return userToken
	}

	userTokenDecrypted := (*userTokenCookie).Value
	encryptionService, err := services.NewEncryptionService()
	userToken, err = encryptionService.Decrypt(userTokenDecrypted)

	if err != nil {
		log.Println("error with decrypting token from cookie: " + err.Error())
	}

	return userToken
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
