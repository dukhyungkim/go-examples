package main

import (
	"bytes"
	"context"
	pb "grpc/proto/helloworld"
	"io"
	"log"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/encoding/protojson"
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
	conn, _ := grpc.DialContext(
		context.Background(),
		"bufnet",
		grpc.WithBlock(),
		grpc.WithContextDialer(bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	gwmux := runtime.NewServeMux()
	if err := pb.RegisterGreeterHandler(context.Background(), gwmux, conn); err != nil {
		log.Fatalf("failed to add handler")
	}
	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	go func() {
		log.Fatalln(gwServer.ListenAndServe())
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestSayHello(t *testing.T) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	tests := []struct {
		name string
		want string
	}{
		{
			name: "world",
			want: "world world",
		},
		{
			name: "123",
			want: "123 world",
		},
	}

	for _, tt := range tests {
		pbytes, _ := protojson.Marshal(&pb.HelloRequest{Name: tt.name})
		buff := bytes.NewBuffer(pbytes)
		resp, err := client.Post("http://localhost:8090/v1/example/echo", "application/json", buff)
		if err != nil {
			t.Errorf("HelloTest(%v) got unexpected error", tt.name)
		}
		assert.Equal(t, resp.StatusCode, 200)
		w, _ := protojson.Marshal(&pb.HelloReply{Message: tt.want})
		respBody, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)
		assert.Equal(t, string(w), string(respBody))
		err = resp.Body.Close()
		if err != nil {
			log.Println(err)
		}
	}
}
