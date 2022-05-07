package main

import (
	"context"
	"fmt"
	"go-examples/common/config"
	pb "go-examples/proto/helloworld"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct {
	pb.GreeterServer
}

func (s *server) SayHello(_ context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	opts, err := config.ParseServerFlags()
	if err != nil {
		log.Panicln(err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", opts.GRPCPort))
	if err != nil {
		log.Panicf("failed to listen: %v", err)
	}

	var s *grpc.Server
	if opts.Cert != "" && opts.PrivateKey != "" {
		var grpcOpts []grpc.ServerOption
		creds, err := credentials.NewServerTLSFromFile(opts.Cert, opts.PrivateKey)
		if err != nil {
			log.Panicln(err)
		}
		grpcOpts = append(grpcOpts, grpc.Creds(creds))
		s = grpc.NewServer(grpcOpts...)
	} else {
		s = grpc.NewServer()
	}

	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	err = s.Serve(lis)
	if err != nil {
		log.Panicf("failed to serve: %v", err)
	}
}
