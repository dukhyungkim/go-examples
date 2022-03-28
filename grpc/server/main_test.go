package main

import (
	"context"
	pb "go-examples/proto/helloworld"
	"log"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestSayHello(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	tests := []struct {
		name string
		want string
	}{
		{
			name: "world",
			want: "Hello world",
		},
		{
			name: "123",
			want: "Hello 123",
		},
	}

	for _, tt := range tests {
		req := &pb.HelloRequest{Name: tt.name}
		resp, err := client.SayHello(context.Background(), req)
		if err != nil {
			t.Errorf("HelloTest(%v) got unexpected error", tt.name)
		}
		assert.Equal(t, tt.want, resp.Message)
	}
}
