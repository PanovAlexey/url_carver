package services

import (
	"context"
	"fmt"
	"net/http"
)

type key string

const (
	userTokenKey key = "token"
)

type ContextStorageService struct {
}

func GetContextStorageService() ContextStorageService {
	return ContextStorageService{}
}

func (service ContextStorageService) SaveUserTokenToContext(r http.Request, userToken string) http.Request {
	ctx := context.WithValue(r.Context(), userTokenKey, userToken)

	return *r.WithContext(ctx)
}

func (service ContextStorageService) GetUserTokenFromContext(ctx context.Context) string {
	return fmt.Sprintf("%v", ctx.Value(userTokenKey))
}
