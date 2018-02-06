package main

import (
	"context"
	"log"
	"net"
	"sync"

	"github.com/nihei9/go-misc/grpc/simple/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	lis, err := net.Listen("tcp", ":10080")
	if err != nil {
		log.Fatalf("failed to listen; error: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterMessageProcessorServer(grpcServer, NewMessageProcessor())
	grpcServer.Serve(lis)
}

type MessageProcessor struct {
	msgChan chan *pb.Message
	wg      *sync.WaitGroup
}

func NewMessageProcessor() *MessageProcessor {
	return &MessageProcessor{
		msgChan: make(chan *pb.Message, 1000),
		wg:      &sync.WaitGroup{},
	}
}

func (p *MessageProcessor) Do(ctx context.Context, msg *pb.Message) (*pb.Result, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md = map[string][]string{}
	}

	log.Printf("Got message; ID: %v, data: \"%v\", metadata: %v", msg.GetId(), msg.GetData(), md)

	return &pb.Result{
		Status:  pb.Result_OK,
		Message: msg.GetId(),
	}, nil
}
