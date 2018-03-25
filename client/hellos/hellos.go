package main

import (
	"api/core"
	"api/greeter"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	address = "secure.front:8001"
)

func main() {
	creds, _ := credentials.NewClientTLSFromFile("/data/cert/pem/crt.pem", "hello")
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := greeter.NewGreeterClient(conn)

	resp, err := client.Say(context.Background(), &greeter.Hello{Query: &core.Info{
		Name: "foo",
	}})

	if err != nil {
		log.Fatalf("fail to say: %v", err)
	}

	println("secure")
	println("name: ", resp.Name.Value)
	println("address: ", resp.PrimaryAddress.Value)
}
