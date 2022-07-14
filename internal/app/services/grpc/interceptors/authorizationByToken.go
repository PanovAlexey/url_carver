package interceptors

import (
	"context"
	"github.com/PanovAlexey/url_carver/config"
	"github.com/PanovAlexey/url_carver/internal/app/services"
	"github.com/PanovAlexey/url_carver/internal/app/services/encryption"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
)

func AuthorizationByToken(
	ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
) (interface{}, error) {
	encryptionService, _ := encryption.NewEncryptionService(config.New())
	userTokenAuthorizationService := services.GetUserTokenAuthorizationService()
	userToken := userTokenAuthorizationService.GetUserTokenFromGRpcMeta(ctx, encryptionService)

	if !userTokenAuthorizationService.IsUserTokenValid(userToken) {
		userToken = userTokenAuthorizationService.UserTokenGenerate()
	}

	userTokenEncrypted := encryptionService.Encrypt(userToken)
	tokenHeader := metadata.New(map[string]string{services.UserTokenName: userTokenEncrypted})

	if err := grpc.SendHeader(ctx, tokenHeader); err != nil { // For using user token on the client.
		log.Println("setting token to gRPC meta error: " + err.Error())
		// @ToDo: добавить нормальный возврат ошибки
	}

	// For using user token on the server down the stack.
	ctx = context.WithValue(ctx, services.UserTokenName, userTokenEncrypted)

	return handler(ctx, req)
}
