package main

import (
	"context"
	"log"

	"github.com/nihei9/go-misc/grpc/server_side_version_up/v2/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:10080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial; error: %v", err)
	}
	defer conn.Close()
	client := pb.NewBookshelfClient(conn)

	book, err := client.Lookup(context.Background(), &pb.LookupRequest{
		Title: "title-1",
	})
	if err != nil {
		log.Fatalf("failed to call Do; error: %v", err)
	}

	log.Printf("result: %#v", book)
}
