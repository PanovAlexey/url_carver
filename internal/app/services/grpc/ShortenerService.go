package grpc

import (
	"context"
	"fmt"
	"github.com/PanovAlexey/url_carver/internal/app/services"
	"github.com/PanovAlexey/url_carver/internal/app/services/database"
	pb "github.com/PanovAlexey/url_carver/pkg/shortener_grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ShortenerService struct {
	pb.UnimplementedShortenerServer

	errorService                  services.ErrorService
	memoryService                 services.MemoryService
	storageService                services.StorageService
	shorteningService             services.ShorteningService
	contextStorageService         services.ContextStorageService
	userTokenAuthorizationService services.UserTokenAuthorizationService
	URLMappingService             services.MappingService
	databaseService               database.DatabaseInterface
	databaseURLService            services.DatabaseURLService
	databaseUserService           services.DatabaseUserService
}

func GetGRPCShortenerService(
	errorService services.ErrorService,
	memoryService services.MemoryService,
	storageService services.StorageService,
	shorteningService services.ShorteningService,
	contextStorageService services.ContextStorageService,
	userTokenAuthorizationService services.UserTokenAuthorizationService,
	databaseService database.DatabaseInterface,
	databaseURLService services.DatabaseURLService,
	databaseUserService services.DatabaseUserService,
) ShortenerService {
	return ShortenerService{
		errorService:                  errorService,
		memoryService:                 memoryService,
		storageService:                storageService,
		shorteningService:             shorteningService,
		contextStorageService:         contextStorageService,
		userTokenAuthorizationService: userTokenAuthorizationService,
		URLMappingService:             services.MappingService{},
		databaseService:               databaseService,
		databaseURLService:            databaseURLService,
		databaseUserService:           databaseUserService,
	}
}

func (s ShortenerService) AddURL(ctx context.Context, request *pb.AddURLRequest) (*pb.AddURLResponse, error) {
	var response pb.AddURLResponse

	if len(request.LongURL) == 0 {
		response.Error = "URL is empty"
		return &response, status.Errorf(codes.InvalidArgument, response.Error, request.LongURL)
	}

	url, err := s.shorteningService.GetURLEntityByLongURL(request.LongURL)

	if err != nil {
		response.Error = err.Error()
		return &response, status.Errorf(codes.InvalidArgument, err.Error(), request.LongURL)
	}

	url.UserID = s.contextStorageService.GetUserTokenFromContext(ctx)
	s.memoryService.SaveURL(url)
	s.storageService.SaveURL(url)
	_, err = s.databaseURLService.SaveURL(url)

	if err != nil || len(url.LongURL) == 0 {
		response.Error = err.Error()

		if s.errorService.IsKeyDuplicated(err) {
			return &response, status.Errorf(codes.AlreadyExists, err.Error(), request.LongURL)
		}

		return &response, status.Errorf(codes.Unknown, err.Error(), request.LongURL)
	}

	shortURLJSON := s.memoryService.GetShortURLDtoByURL(url)

	fmt.Println("URL " + url.LongURL + " added by " + url.ShortURL)
	response.ShortURL = shortURLJSON.Value

	return &response, nil
}

func (s ShortenerService) GetURLByShort(ctx context.Context, in *pb.GetURLRequest) (*pb.GetURLResponse, error) {
	var response pb.GetURLResponse
	response.LongURL = "longURLMock"
	response.Error = ""

	return &response, nil
}

func (s ShortenerService) AddBatchURLs(ctx context.Context, in *pb.AddBatchURLRequest) (*pb.AddBatchURLResponse, error) {
	var response pb.AddBatchURLResponse
	var batchURLItem pb.BatchURLItem

	batchURLItem.LongURL = "longURL01Mock"
	batchURLItem.CorrelationID = "1"

	response.BatchURL = append(response.BatchURL, &batchURLItem)
	response.Error = ""

	return &response, nil
}

func (s ShortenerService) GetURLsByUser(ctx context.Context, in *emptypb.Empty) (*pb.GetURLsByUserResponse, error) {
	var response pb.GetURLsByUserResponse
	var URLComplex pb.URLComplex

	URLComplex.ShortURL = "shortURLMock"
	URLComplex.OriginalURL = "longURLMock"

	response.URLs = append(response.URLs, &URLComplex)
	response.Error = ""

	return &response, nil
}

func (s ShortenerService) DeleteURLs(ctx context.Context, in *pb.DeleteURLsRequest) (*emptypb.Empty, error) {
	return nil, nil
}

func (s ShortenerService) GetStats(ctx context.Context, in *emptypb.Empty) (*pb.GetStatsResponse, error) {
	var response pb.GetStatsResponse
	response.Users = 123
	response.URLS = 234

	return &response, nil
}
