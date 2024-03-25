package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/Dev9tkov/go_chat/pkg/chat_v1"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedChatV1Server
}

func (s *server) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	fmt.Printf("Req %s", req.GetUsernames())
	return &desc.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func (s *server) Delete(_ context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	fmt.Printf("Delete %d", req.GetId())
	return nil, nil
}

func (s *server) SendMessage(_ context.Context, req *desc.SendRequest) (*emptypb.Empty, error) {
	fmt.Printf("Send %s", req.From)
	return nil, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
