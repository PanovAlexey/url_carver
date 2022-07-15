package services

import (
	"context"
	"github.com/PanovAlexey/url_carver/internal/app/services/encryption"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
	"log"
	"net/http"
	"time"
)

type key string

const UserTokenName key = `token`
const UserTokenCookieExpirationDate = 12 * time.Hour

type UserTokenAuthorizationService struct {
}

func GetUserTokenAuthorizationService() *UserTokenAuthorizationService {
	return &UserTokenAuthorizationService{}
}

func (service UserTokenAuthorizationService) IsUserTokenValid(userToken string) bool {
	return len(userToken) >= 1
}

func (service UserTokenAuthorizationService) UserTokenGenerate() string {
	return uuid.New().String()
}

func (service UserTokenAuthorizationService) GetUserTokenFromCookie(
	r *http.Request, encryptionService encryption.EncryptorInterface,
) string {
	userToken := ``
	userTokenCookie, err := r.Cookie(string(UserTokenName))

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

func (service UserTokenAuthorizationService) GetUserTokenFromGRpcMeta(
	ctx context.Context, encryptionService encryption.EncryptorInterface,
) string {
	userTokenEncrypted := ``

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		values := md.Get(string(UserTokenName))

		if len(values) != 0 {
			userTokenEncrypted = values[0]
		}
	}

	if len(userTokenEncrypted) == 0 {
		log.Println("error with getting token from gRPC metadata")

		return ""
	}

	userTokenDecrypted, err := encryptionService.Decrypt(userTokenEncrypted)

	if err != nil {
		log.Println("error with decrypting token from gRPC metadata: " + err.Error())

		return ""
	}

	return userTokenDecrypted
}

func (service UserTokenAuthorizationService) SetUserTokenToCookie(userToken string, w http.ResponseWriter) {
	cookie := http.Cookie{
		Name:    string(UserTokenName),
		Value:   userToken,
		Expires: time.Now().Add(UserTokenCookieExpirationDate),
		Path:    `/`,
	}

	http.SetCookie(w, &cookie)
}
