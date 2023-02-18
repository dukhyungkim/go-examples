package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"grpc/config"
	pb "grpc/proto/helloworld"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	opts, err := config.ParseClientFlags()
	if err != nil {
		log.Panicln(err)
	}

	var conn *grpc.ClientConn

	if opts.Cert != "" {
		tlsCredentials, err := loadTLSCredentials(opts.Cert)
		if err != nil {
			log.Panicln("cannot load TLS credentials: ", err)
		}

		conn, err = grpc.Dial(opts.Target, grpc.WithTransportCredentials(tlsCredentials), grpc.WithBlock())
		if err != nil {
			log.Panicf("did not connect: %v", err)
		}
	} else {
		conn, err = grpc.Dial(opts.Target, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
		if err != nil {
			log.Panicf("did not connect: %v", err)
		}
	}
	defer func() {
		if err = conn.Close(); err != nil {
			log.Println(err)
		}
	}()
	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: opts.Name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}

func loadTLSCredentials(cert string) (credentials.TransportCredentials, error) {
	pemServerCA, err := os.ReadFile(cert)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	cfg := &tls.Config{
		RootCAs: certPool,
	}

	return credentials.NewTLS(cfg), nil
}
