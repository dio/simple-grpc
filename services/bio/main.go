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
	address = "0.0.0.0:3002"
)

type server struct{}

func (s *server) GetName(ctx context.Context, in *core.Info) (*people.Name, error) {
	return &people.Name{Value: "ok"}, nil
}

func main() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	people.RegisterBioServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
