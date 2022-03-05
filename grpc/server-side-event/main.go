package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	pb "go-examples/grpc/server-side-event/sse"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panicln(err)
	}
	defer func() {
		if err = lis.Close(); err != nil {
			log.Println(err)
		}
	}()

	s := grpc.NewServer()
	pb.RegisterStreamServiceServer(s, &server{})

	log.Println("start grpc server")
	go func() {
		log.Panicln(s.Serve(lis))
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Panicln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = pb.RegisterStreamServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Panicln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: cors.AllowAll().Handler(gwmux),
	}

	log.Println("start grpc-gateway server")
	log.Panic(gwServer.ListenAndServe())
}
