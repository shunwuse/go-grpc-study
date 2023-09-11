package main

import (
	"context"
	pb "go-grpc-study/protos"
	"log"
)

type grpcService struct {
	pb.UnimplementedGreeterServer
}

func (s *grpcService) SayHello(ctx context.Context, req *pb.GreeterRequest) (*pb.GreeterResponse, error) {
	log.Printf("Received: %v", req.Name)
	return &pb.GreeterResponse{Message: "Hello " + req.Name}, nil
}
