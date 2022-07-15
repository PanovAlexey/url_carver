package grpc

import (
	"context"
	"github.com/PanovAlexey/url_carver/internal/app/domain/dto"
	"github.com/PanovAlexey/url_carver/internal/app/domain/dto/batch"
	"github.com/PanovAlexey/url_carver/internal/app/domain/entity/url"
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

func (s ShortenerService) AddBatchURLs(ctx context.Context, request *pb.AddBatchURLRequest) (*pb.AddBatchURLResponse, error) {
	var response pb.AddBatchURLResponse

	var URLCollection []url.URL
	batchInputURLDTOCollection := []batch.BatchInputURL{}
	batchOutputURLDTOCollection := []batch.BatchOutputURL{}

	for _, v := range request.BatchURL {
		batchInputURLDTOCollection = append(
			batchInputURLDTOCollection,
			batch.BatchInputURL{CorrelationID: v.CorrelationID, OriginalURL: v.LongURL},
		)
	}

	for _, URL := range request.BatchURL {
		url, err := s.shorteningService.GetURLEntityByLongURL(URL.LongURL)

		if err != nil || len(url.LongURL) == 0 {
			return &response, status.Errorf(
				codes.ResourceExhausted, `error while getting URL entity by long URL: `+URL.LongURL, URL.LongURL,
			)
		}

		url.UserID = s.contextStorageService.GetUserTokenFromContext(ctx)

		s.memoryService.SaveURL(url)
		s.storageService.SaveURL(url)

		shortURLWithDomain := s.shorteningService.GetShortURLWithDomain(url.ShortURL)

		batchOutputURLDTOCollection = append(
			batchOutputURLDTOCollection, batch.NewBatchOutputURL(URL.CorrelationID, shortURLWithDomain),
		)

		response.BatchURL = append(response.BatchURL, &pb.BatchURLItem{LongURL: url.LongURL, CorrelationID: url.ShortURL})
		URLCollection = append(URLCollection, url)
	}

	s.databaseURLService.SaveBatchURLs(URLCollection)

	return &response, nil
}

func (s ShortenerService) GetURLsByUser(ctx context.Context, request *emptypb.Empty) (*pb.GetURLsByUserResponse, error) {
	var response pb.GetURLsByUserResponse

	userToken := s.contextStorageService.GetUserTokenFromContext(ctx)

	if !s.userTokenAuthorizationService.IsUserTokenValid(userToken) {
		return &response, status.Errorf(codes.Unauthenticated, "authorization failed", userToken)
	}

	URLCollection := s.memoryService.GetURLsByUserToken(userToken)

	if len(URLCollection) < 1 {
		return &response, nil
	}

	URLCollectionForShowingUser := s.URLMappingService.MapURLEntityCollectionToDTO(URLCollection)

	var URLComplex pb.URLComplex

	for _, v := range URLCollectionForShowingUser {
		URLComplex.ShortURL = v.ShortURL
		URLComplex.OriginalURL = v.LongURL

		response.URLs = append(response.URLs, &URLComplex)
	}

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
