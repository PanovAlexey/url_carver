package services

import (
	"github.com/PanovAlexey/url_carver/internal/app/services/encryption"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

const UserTokenName = `token`
const UserTokenCookieExpirationDate = 1 * time.Minute

type userTokenAuthorizationService struct {
}

func GetUserTokenAuthorizationService() *userTokenAuthorizationService {
	return &userTokenAuthorizationService{}
}

func (service userTokenAuthorizationService) IsUserTokenValid(userToken string) bool {
	return len(userToken) >= 1
}

func (service userTokenAuthorizationService) UserTokenGenerate() string {
	return uuid.New().String()
}

func (service userTokenAuthorizationService) GetUserTokenFromCookie(
	r *http.Request, encryptionService encryption.EncryptorInterface,
) string {
	userToken := ``
	userTokenCookie, err := r.Cookie(UserTokenName)

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

func (service userTokenAuthorizationService) SetUserTokenToCookie(userToken string, w http.ResponseWriter) {
	cookie := http.Cookie{
		Name:    UserTokenName,
		Value:   userToken,
		Expires: time.Now().Add(UserTokenCookieExpirationDate),
	}

	http.SetCookie(w, &cookie)
}
