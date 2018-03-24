package main

import (
	"context"
	"log"
	"net"

	"api/core"
	"api/greeter"

	"hello/service"

	"google.golang.org/grpc"
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

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
