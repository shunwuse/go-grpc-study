package main

import (
	"context"
	pb "go-grpc-study/protos"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithPerRPCCredentials(&clientTokenAuth{}),
	}

	// connect to server
	conn, err := grpc.Dial(":8080", opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return
	}
	defer conn.Close()

	// create client
	client := pb.NewGreeterClient(conn)

	// call service
	callGrpc(client, "World")
}

func callGrpc(client pb.GreeterClient, name string) {
	// set timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// call service
	resp, err := client.SayHello(ctx, &pb.GreeterRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
		return
	}

	// print response
	log.Printf("Response: %s", resp.Message)
}
