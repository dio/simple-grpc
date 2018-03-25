package main

import (
	"context"
	"log"
	"net"

	"api/core"
	"api/greeter"

	"hello/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

const (
	address  = "0.0.0.0:5001"
	upstream = "internal:10000"
)

type server struct {
	endpoint *service.Endpoint
}

func (s *server) Say(ctx context.Context, in *greeter.Hello) (*greeter.Reply, error) {
	name, err := s.endpoint.GetName(ctx, &core.Info{})
	if err != nil {
		return nil, err
	}
	primary, err := s.endpoint.GetPrimary(ctx, &core.Info{})
	if err != nil {
		return nil, err
	}
	return &greeter.Reply{
		Name:           name,
		PrimaryAddress: primary,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	endpoint, err := service.NewClient(upstream)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	s := grpc.NewServer()
	greeter.RegisterGreeterServer(s, &server{
		endpoint: endpoint,
	})

	hs := health.NewServer()
	hs.SetServingStatus("grpc.health.v1.hello", 1)
	healthpb.RegisterHealthServer(s, hs)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
