package middleware

import (
	"github.com/PanovAlexey/url_carver/internal/app/services"
	"github.com/PanovAlexey/url_carver/internal/app/services/encryption"
	"log"
	"net/http"
	"time"
)

func Authorization(encryptionService encryption.EncryptorInterface) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userToken := getUserTokenFromCookie(r, encryptionService)
			userTokenAuthorizationService := services.GetUserTokenAuthorizationService()

			if !userTokenAuthorizationService.IsUserTokenValid(userToken) {
				userToken = userTokenAuthorizationService.UserTokenGenerate()
				userTokenEncrypted := encryptionService.Encrypt(userToken)
				setUserTokenToCookie(userTokenEncrypted, w)
			}

			next.ServeHTTP(w, r)
		})
	}
}

func getUserTokenFromCookie(r *http.Request, encryptionService encryption.EncryptorInterface) string {
	userToken := ``
	userTokenCookie, err := r.Cookie(services.UserTokenName)

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

func setUserTokenToCookie(userToken string, w http.ResponseWriter) {
	cookie := http.Cookie{
		Name:    services.UserTokenName,
		Value:   userToken,
		Expires: time.Now().Add(services.UserTokenCookieExpirationDate),
	}

	http.SetCookie(w, &cookie)
}
