package middleware

import (
	"github.com/PanovAlexey/url_carver/internal/app/services"
	"github.com/PanovAlexey/url_carver/internal/app/services/encryption"
	"net/http"
)

func Authorization(encryptionService encryption.EncryptorInterface) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userTokenAuthorizationService := services.GetUserTokenAuthorizationService()
			userToken := userTokenAuthorizationService.GetUserTokenFromCookie(r, encryptionService)

			if !userTokenAuthorizationService.IsUserTokenValid(userToken) {
				userToken = userTokenAuthorizationService.UserTokenGenerate()
				userTokenEncrypted := encryptionService.Encrypt(userToken)
				userTokenAuthorizationService.SetUserTokenToCookie(userTokenEncrypted, w)
			}

			contextStorageService := services.GetContextStorageService()
			*r = contextStorageService.SaveUserTokenToContext(*r, userToken)

			next.ServeHTTP(w, r)
		})
	}
}
