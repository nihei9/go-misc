package main

import (
	"context"
	"log"

	"github.com/nihei9/go-misc/grpc/simple/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn, err := grpc.Dial("localhost:10080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial; error: %v", err)
	}
	defer conn.Close()
	client := pb.NewMessageProcessorClient(conn)

	md := metadata.Pairs("foo", "xxxx")
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	result, err := client.Do(ctx, &pb.Message{
		Id:   "24281f72-60b9-4e76-9d0c-17fb3dd3c8c7",
		Data: "Hello.",
	})
	if err != nil {
		log.Fatalf("failed to call Do; error: %v", err)
	}

	log.Printf("result: %v", result)
}
