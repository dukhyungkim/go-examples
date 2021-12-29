package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go-examples/common/config"
	pb "go-examples/proto/order"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

type server struct {
	pb.OrderServiceServer
}

func (s *server) SayOrder(_ context.Context, in *pb.OrderRequest) (*pb.OrderReply, error) {
	log.Println("Received:", in.GetName(), in.GetContext(), in.GetAge())
	return &pb.OrderReply{Name: in.GetName(), Context: in.GetContext(), Age: in.GetAge()}, nil
}

func main() {
	opts, err := config.ParseServerFlags()
	if err != nil {
		log.Fatalln(err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", opts.GRPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterOrderServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		fmt.Sprintf("0.0.0.0:%d", opts.GRPCPort),
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = pb.RegisterOrderServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", opts.HTTPPort),
		Handler: gwmux,
	}

	log.Printf("Serving gRPC-Gateway on http://0.0.0.0:%d\n", opts.HTTPPort)
	log.Fatalln(gwServer.ListenAndServe())
}
