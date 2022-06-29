package services

import (
	"fmt"
	"github.com/PanovAlexey/url_carver/config"
	"github.com/PanovAlexey/url_carver/internal/app/services/encryption"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_GetUserTokenAuthorizationService(t *testing.T) {
	t.Run("Test user token authorization service creating", func(t *testing.T) {
		userTokenAuthorizationService := GetUserTokenAuthorizationService()
		structType := fmt.Sprintf("%T", userTokenAuthorizationService)

		assert.Equal(t, structType, "*services.UserTokenAuthorizationService")
	})
}

func Test_IsUserTokenValid(t *testing.T) {
	t.Run("Test is user token validation", func(t *testing.T) {
		userTokenAuthorizationService := GetUserTokenAuthorizationService()

		assert.Equal(t, false, userTokenAuthorizationService.IsUserTokenValid(""))
	})
}

func Test_UserTokenGenerate(t *testing.T) {
	t.Run("Test user token generating", func(t *testing.T) {
		userTokenAuthorizationService := GetUserTokenAuthorizationService()

		assert.Equal(t, 36, len(userTokenAuthorizationService.UserTokenGenerate()))
	})
}

func Test_GetUserTokenFromCookie(t *testing.T) {
	t.Run("Test getting user token from cookie", func(t *testing.T) {
		userTokenAuthorizationService := GetUserTokenAuthorizationService()

		r, _ := http.NewRequest(http.MethodGet, "www.example.com", nil)
		encryptionService, _ := encryption.NewEncryptionService(config.New())

		cookieCorrect := http.Cookie{
			Name:     UserTokenName,
			Value:    "d8270dd3786a2b6fd5594d427a67c9220aaea59797017b671907ff36c9fdab4f215c51651b37e428db3a672a59509283393a1f16",
			HttpOnly: false,
		}

		r.AddCookie(&cookieCorrect)
		assert.Equal(t, 36, len(userTokenAuthorizationService.GetUserTokenFromCookie(r, encryptionService)))
	})
}

func Test_GetUserTokenFromCookieWrong(t *testing.T) {
	t.Run("Test getting user token from cookie", func(t *testing.T) {
		userTokenAuthorizationService := GetUserTokenAuthorizationService()

		r, _ := http.NewRequest(http.MethodGet, "www.example.com", nil)
		encryptionService, _ := encryption.NewEncryptionService(config.New())
		assert.Equal(t, 0, len(userTokenAuthorizationService.GetUserTokenFromCookie(r, encryptionService)))

		cookie := http.Cookie{
			Name:     UserTokenName,
			Value:    "123456",
			HttpOnly: false,
		}
		r.AddCookie(&cookie)
		assert.Equal(t, 0, len(userTokenAuthorizationService.GetUserTokenFromCookie(r, encryptionService)))
	})
}

type ResponseWriterTest struct {
}

func (responseWriterTest ResponseWriterTest) Header() http.Header {
	return map[string][]string{}
}

func (responseWriterTest ResponseWriterTest) Write([]byte) (int, error) {
	return 0, nil
}

func (responseWriterTest ResponseWriterTest) WriteHeader(statusCode int) {
}

// @ToDo: finishing this test
func Test_SetUserTokenToCookie(t *testing.T) {
	t.Run("Test set user token to cookie", func(t *testing.T) {
		userTokenAuthorizationService := GetUserTokenAuthorizationService()
		userToken := "test_token"

		responseWriterTest := ResponseWriterTest{}
		userTokenAuthorizationService.SetUserTokenToCookie(userToken, responseWriterTest)

	})
}
