package main

import (
	"context"
	"log"
	"net"

	"github.com/nihei9/go-misc/grpc/server_side_version_up/v1/pb"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":10080")
	if err != nil {
		log.Fatalf("failed to listen; error: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterBookshelfServer(grpcServer, NewBookshelf())
	grpcServer.Serve(lis)
}

type book struct {
	title  string
	auther string
}

type Bookshelf struct {
	records []*book
}

func NewBookshelf() *Bookshelf {
	return &Bookshelf{
		records: []*book{
			&book{
				title:  "title-0",
				auther: "auther-0",
			},
			&book{
				title:  "title-1",
				auther: "auther-1",
			},
			&book{
				title:  "title-2",
				auther: "auther-2",
			},
		},
	}
}

func (b *Bookshelf) Lookup(ctx context.Context, req *pb.LookupRequest) (*pb.BookRecord, error) {
	log.Printf("Got a message; %#v", req)

	book := &pb.BookRecord{}
	for _, record := range b.records {
		if record.title == req.GetTitle() {
			book = &pb.BookRecord{
				Title:  record.title,
				Auther: record.auther,
			}
		}
	}

	return book, nil
}
