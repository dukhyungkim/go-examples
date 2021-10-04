package main

import (
	"context"
	"fmt"
	"go-examples/common/config"
	pb "go-examples/grpc/helloworld"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.GreeterServer
}

func (s *server) SayHello(_ context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	opts, err := config.ParseRPCFlags()
	if err != nil {
		log.Fatalln(err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", opts.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
