package grpc

import (
	"context"
	pb "github.com/PanovAlexey/url_carver/pkg/shortener_grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ShortenerService struct {
	pb.UnimplementedShortenerServer
}

func (s ShortenerService) AddURL(ctx context.Context, in *pb.AddURLRequest) (*pb.AddURLResponse, error) {
	var response pb.AddURLResponse
	response.ShortURL = "shortURLMock"
	response.Error = ""

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
