package main

import (
	"context"
	"errors"
	pb "go-grpc-study/protos"
	"log"

	"google.golang.org/grpc/metadata"
)

type grpcService struct {
	pb.UnimplementedGreeterServer
}

func (s *grpcService) SayHello(ctx context.Context, req *pb.GreeterRequest) (*pb.GreeterResponse, error) {
	// get token from context
	metadata, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("missing context metadata")
	}

	appId := metadata["appid"]
	appKey := metadata["appkey"]
	log.Printf("appid: %s, appkey: %s", appId, appKey)

	if appId[0] != "myApp" || appKey[0] != "myAppKey" {
		return nil, errors.New("invalid token")
	}

	log.Printf("Received: %v", req.Name)
	return &pb.GreeterResponse{Message: "Hello " + req.Name}, nil
}
