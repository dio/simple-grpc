package main

import (
	"api/core"
	"api/people"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

const (
	address = "0.0.0.0:3001"
)

var (
	data = map[string]string{
		"foo": "bar",
	}
)

type server struct{}

func (s *server) GetPrimaryAddress(ctx context.Context, in *core.Info) (*people.PrimaryAddress, error) {
	name := data[in.Name]
	if name == "" {
		name = "bar"
	}
	return &people.PrimaryAddress{Value: name}, nil
}

func main() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	people.RegisterAddressServer(s, &server{})

	hs := health.NewServer()
	hs.SetServingStatus("grpc.health.v1.address", 1)
	healthpb.RegisterHealthServer(s, hs)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
