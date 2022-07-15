package grpc

import (
	"context"
	"github.com/PanovAlexey/url_carver/internal/app/domain/dto"
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

	response.ShortURL = url.ShortURL

	return &response, nil
}

func (s ShortenerService) GetURLByShort(ctx context.Context, request *pb.GetURLRequest) (*pb.GetURLResponse, error) {
	var response pb.GetURLResponse

	if len(request.ShortURL) == 0 || !s.memoryService.IsExistURLEntityByShortURL(request.ShortURL) {
		response.Error = "Not found"
		return &response, status.Errorf(codes.NotFound, response.Error, request.ShortURL)
	}

	urlFull, err := s.memoryService.GetURLEntityByShortURL(request.ShortURL)

	if err != nil {
		if s.errorService.IsDeleted(err) {
			response.Error = "Deleted"
			return &response, status.Errorf(codes.ResourceExhausted, response.Error, request.ShortURL)
		} else {
			response.Error = "NotFound"
			return &response, status.Errorf(codes.NotFound, response.Error, request.ShortURL)
		}
	}

	response.LongURL = urlFull

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

func (s ShortenerService) DeleteURLs(ctx context.Context, request *pb.DeleteURLsRequest) (*emptypb.Empty, error) {
	userToken := s.contextStorageService.GetUserTokenFromContext(ctx)

	if !s.userTokenAuthorizationService.IsUserTokenValid(userToken) {
		return &emptypb.Empty{}, status.Errorf(codes.NotFound, "authorization failed", userToken)
	}

	err := s.databaseURLService.RemoveByShortURLSlice(request.ShortURL, userToken)
	s.memoryService.DeleteURLsByShortValueSlice(request.ShortURL)

	if err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Unknown, "URL deletion failed: "+err.Error(), userToken)
	}

	return &emptypb.Empty{}, nil
}

func (s ShortenerService) GetStats(ctx context.Context, in *emptypb.Empty) (*pb.GetStatsResponse, error) {
	var response pb.GetStatsResponse
	usersCount, err := s.databaseUserService.GetAllUsersCount()

	if err != nil {
		return &response, status.Errorf(codes.Unknown, err.Error())
	}

	appStat := dto.GetAppStatByURLsCountAndUsersCount(s.memoryService.GetAllURLsCount(), usersCount)
	response.URLS = int32(appStat.Urls)
	response.Users = int32(appStat.Users)

	return &response, nil
}
