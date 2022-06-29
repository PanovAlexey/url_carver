package services

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_GetUserTokenFromContext(t *testing.T) {
	t.Run("Test getting user token from context", func(t *testing.T) {
		contextStorageService := GetContextStorageService()
		structType := fmt.Sprintf("%T", contextStorageService)

		assert.Equal(t, structType, "services.ContextStorageService")
	})
}

func Test_SaveUserTokenToContextAndGetUserTokenFromContext(t *testing.T) {
	t.Run("Test saving user token to context and getting user token from context", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "www.example.com", nil)
		sourceUserToken := "test_user_token_value"
		contextStorageService := GetContextStorageService()
		resultRequest := contextStorageService.SaveUserTokenToContext(*request, sourceUserToken)
		resultUserToken := contextStorageService.GetUserTokenFromContext(resultRequest.Context())

		assert.Equal(t, sourceUserToken, resultUserToken)
	})
}
