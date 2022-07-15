package main

import (
	"context"
	"github.com/PanovAlexey/url_carver/config"
	"github.com/PanovAlexey/url_carver/internal/app/domain/dto"
	"github.com/PanovAlexey/url_carver/internal/app/domain/entity/url"
	pb "github.com/PanovAlexey/url_carver/pkg/shortener_grpc"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"testing"
)

func TestGRpcURLManaging(t *testing.T) {
	conn := getConn()
	defer conn.Close()

	correctURL := url.URL{
		LongURL:  "http://grpc3.com",
		ShortURL: "e5a6c6bd9177a0bf975f9331805593cc",
	}
	userTokenSource := "8d210c867a31743ed50b4d427a67c97f0bfaa5cf960328674e56fd37cdf3f8472d5d0d6419fce9db74d72d803e35c71ad403156c"

	appStatBeforeTests, token, err := getStats(conn, userTokenSource)

	if err != nil {
		return
	}

	assert.Equal(t, nil, err)
	assert.Equal(t, userTokenSource, token)

	shortURL, token, err := addURL(conn, correctURL.LongURL, userTokenSource)

	assert.Equal(t, nil, err)
	assert.Equal(t, correctURL.ShortURL, shortURL)
	assert.NotEmpty(t, token)

	longURL, token, err := getURL(conn, correctURL.ShortURL, userTokenSource)

	assert.Equal(t, nil, err)
	assert.Equal(t, correctURL.LongURL, longURL)
	assert.Equal(t, userTokenSource, token)

	token, err = deleteURL(conn, correctURL.ShortURL, userTokenSource)
	assert.Equal(t, nil, err)
	assert.Equal(t, userTokenSource, token)

	appStatAfterTests, token, err := getStats(conn, userTokenSource)
	assert.Equal(t, nil, err)
	assert.Equal(t, appStatBeforeTests.Users+1, appStatAfterTests.Users)
	assert.Equal(t, userTokenSource, token)

	userURLs, token, err := getURLsByUser(conn, userTokenSource)
	assert.Equal(t, nil, err)
	assert.Equal(t, userTokenSource, token)
	assert.Equal(t, correctURL.LongURL, userURLs[0].LongURL)

	var batchURLs []url.URL

	batchURLs = append(batchURLs, url.URL{
		LongURL:  "https://github.com",
		ShortURL: "e5a6c6bd9177a0bf975f9331805593cc",
	})

	batchURLs = append(batchURLs, url.URL{
		LongURL:  "https://www.amazon.com",
		ShortURL: "e5a6c6bd9177a0bf975f9331805593cc",
	})

	userTokenSource = "8e7108832f3d7c3ad55f4f437967c92b5dfaa5ccc85228674f02a936cef5f64522085d661c67b1ba18602a03bd9a67e3eca2c463"

	batchURLForShowingUser, token, err := addBatch(conn, batchURLs, userTokenSource)

	assert.Equal(t, nil, err)
	assert.Equal(t, userTokenSource, token)
	assert.Equal(t, 2, len(batchURLForShowingUser))
}

func getConn() *grpc.ClientConn {
	config := config.New()
	conn, err := grpc.Dial("localhost:"+config.GetGRpcServerPort(), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	return conn
}

func getURLsByUser(conn *grpc.ClientConn, inputToken string) (userURLs []dto.URLForShowingUser, outputToken string, err error) {
	var header metadata.MD
	client := pb.NewShortenerClient(conn)
	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(ctx, "token", inputToken)

	response, err := client.GetURLsByUser(ctx, &emptypb.Empty{}, grpc.Header(&header))
	outputToken = header.Get("token")[0]

	for _, v := range response.URLs {
		userURLs = append(userURLs, dto.URLForShowingUser{ShortURL: v.ShortURL, LongURL: v.OriginalURL})
	}

	return
}

func addURL(conn *grpc.ClientConn, longURL string, inputToken string) (shortURL, outputToken string, err error) {
	var header metadata.MD

	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(ctx, "token", inputToken)
	client := pb.NewShortenerClient(conn)
	response, err := client.AddURL(ctx, &pb.AddURLRequest{LongURL: longURL}, grpc.Header(&header))

	if response != nil {
		shortURL = response.ShortURL
	}

	outputToken = header.Get("token")[0]

	return
}

func getURL(conn *grpc.ClientConn, shortURL, inputToken string) (longURL, outputToken string, err error) {
	var header metadata.MD

	client := pb.NewShortenerClient(conn)
	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(ctx, "token", inputToken)
	response, err := client.GetURLByShort(
		ctx, &pb.GetURLRequest{ShortURL: shortURL},
		grpc.Header(&header),
	)

	if response != nil {
		longURL = response.LongURL
	}

	outputToken = header.Get("token")[0]

	return
}

func deleteURL(conn *grpc.ClientConn, shortURL, inputToken string) (outputToken string, err error) {
	var header metadata.MD

	client := pb.NewShortenerClient(conn)
	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(ctx, "token", inputToken)

	_, err = client.DeleteURLs(ctx, &pb.DeleteURLsRequest{ShortURL: []string{shortURL}}, grpc.Header(&header))
	outputToken = header.Get("token")[0]

	return
}

func getStats(conn *grpc.ClientConn, inputToken string) (appStat dto.AppStat, outputToken string, err error) {
	var header metadata.MD
	client := pb.NewShortenerClient(conn)
	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(ctx, "token", inputToken)

	response, err := client.GetStats(ctx, &emptypb.Empty{}, grpc.Header(&header))
	if err == nil {
		appStat.Users = int(response.Users)
		appStat.Urls = int(response.URLS)

		outputToken = header.Get("token")[0]
	}

	return
}

func addBatch(conn *grpc.ClientConn, batchURL []url.URL, inputToken string) (batchURLForShowingUser []dto.URLForShowingUser, outputToken string, err error) {
	var header metadata.MD

	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(ctx, "token", inputToken)
	client := pb.NewShortenerClient(conn)

	var batchURLRequest pb.AddBatchURLRequest

	for _, v := range batchURL {
		batchURLRequest.BatchURL = append(batchURLRequest.BatchURL, &pb.BatchURLItem{CorrelationID: v.ShortURL, LongURL: v.LongURL})
	}

	response, err := client.AddBatchURLs(ctx, &batchURLRequest, grpc.Header(&header))

	for _, v := range response.BatchURL {
		batchURLForShowingUser = append(batchURLForShowingUser, dto.NewURLForShowingUser(v.LongURL, v.CorrelationID, inputToken))
	}

	outputToken = header.Get("token")[0]

	return
}
