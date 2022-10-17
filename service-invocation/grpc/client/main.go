package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/metadata"
)

const (
	address = "localhost:50001"
	appID   = "grpc-server"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	ctx := metadata.AppendToOutgoingContext(context.TODO(), "dapr-app-id", appID)

	for {
		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "Bob"})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}

		log.Printf("res: %s", r.GetMessage())

		time.Sleep(1 * time.Second)
	}
}
