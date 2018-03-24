package main

import (
	"api/core"
	"api/greeter"
	"context"
	"log"

	"google.golang.org/grpc"
)

const (
	address = "plain.front:8000"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := greeter.NewGreeterClient(conn)

	resp, err := client.Say(context.Background(), &greeter.Hello{Query: &core.Info{
		Name: "OK",
	}})

	if err != nil {
		log.Fatalf("fail to say: %v", err)
	}
	println(resp.Name.Value)
	println(resp.PrimaryAddress.Value)
}
