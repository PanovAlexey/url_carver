package middleware

import (
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

const userTokenName = `token`
const userTokenCookieExpirationDate = 1 * time.Minute

type encryptorInterface interface {
	Encrypt(data string) string
	Decrypt(encryptedData string) (string, error)
}

func Authorization(encryptionService encryptorInterface) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userToken := getUserTokenFromCookie(r, encryptionService)

			if !isUserTokenValid(userToken) {
				userToken = userTokenGenerate()
				userTokenEncrypted := encryptionService.Encrypt(userToken)
				setUserTokenToCookie(userTokenEncrypted, w)
			}

			next.ServeHTTP(w, r)
		})
	}
}

func getUserTokenFromCookie(r *http.Request, encryptionService encryptorInterface) string {
	userToken := ``
	userTokenCookie, err := r.Cookie(userTokenName)

	if err != nil {
		if err != http.ErrNoCookie {
			log.Println("error with getting token from cookie: " + err.Error())
		}

		return userToken
	}

	userTokenEncrypted := (*userTokenCookie).Value
	userToken, err = encryptionService.Decrypt(userTokenEncrypted)

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
