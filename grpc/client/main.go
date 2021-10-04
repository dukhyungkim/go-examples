package main

import (
	"context"
	"go-examples/common/config"
	pb "go-examples/grpc/helloworld"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	opts, err := config.ParseClientFlags()
	if err != nil {
		log.Fatalln(err)
	}

	conn, err := grpc.Dial(opts.Target, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: opts.Name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
