package services

import (
	"github.com/google/uuid"
	"time"
)

const UserTokenName = `token`
const UserTokenCookieExpirationDate = 1 * time.Minute

type userTokenAuthorizationService struct {
}

func GetUserTokenAuthorizationService() *userTokenAuthorizationService {
	return &userTokenAuthorizationService{}
}

func (userTokenAuthorizationService userTokenAuthorizationService) IsUserTokenValid(userToken string) bool {
	return len(userToken) >= 1
}

func (userTokenAuthorizationService userTokenAuthorizationService) UserTokenGenerate() string {
	return uuid.New().String()
}
