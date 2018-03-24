package main

import (
	"api/core"
	"api/people"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	address = "0.0.0.0:3001"
)

type server struct{}

func (s *server) GetPrimaryAddress(ctx context.Context, in *core.Info) (*people.PrimaryAddress, error) {
	return &people.PrimaryAddress{Value: "ok"}, nil
}

func main() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	people.RegisterAddressServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
