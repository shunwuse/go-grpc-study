package main

import (
	"log"
	"net"

	pb "go-grpc-study/protos"

	"google.golang.org/grpc"
)

const port = ":8080"

func main() {
	// print message
	log.Println("Starting server...")

	// listen on port 8080
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create grpc service
	grpcServer := grpc.NewServer()
	log.Println("gRPC server is running.")

	// register service
	pb.RegisterGreeterServer(grpcServer, &grpcService{})

	// start service
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
		return
	}
}
