package services

import (
	"context"
	"fmt"
	"net/http"
)

const userTokenKey = "token"

type contextStorageService struct {
}

func GetContextStorageService() contextStorageService {
	return contextStorageService{}
}

func (service contextStorageService) SaveUserTokenToContext(r http.Request, userToken string) http.Request {
	ctx := context.WithValue(r.Context(), userTokenKey, userToken)
	return *r.WithContext(ctx)
}

func (service contextStorageService) GetUserIdFromContext(ctx context.Context) string {
	return fmt.Sprintf("%v", ctx.Value(userTokenKey))
}
