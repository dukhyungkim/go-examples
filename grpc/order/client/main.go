package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"go-examples/common/config"
	pb "go-examples/proto/order"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
	"io/ioutil"
	"log"
	"time"
)

const jsonData = `{
  "master_id": "V40)aVQ.U67#3!(F",
  "context": {
    "PCB_ID": "V40)aVQ.U67#3!(F",
    "ORDER_ID": "576479",
    "FACTORY": "200",
    "MAT_ID": "A6M^C#^j3*T#597S",
    "MAT_CMF_1": "4",
    "END_RES_ID": "Q9366G!I&*S%Xh2^,&15N86aG^HP9.X@^,*Z88M6(Xz^#J6X^6",
    "LOT_ID": "6LGCU0&31*8@)&My,93#&w7*9CH7M&#KQ,%)Q$n2&5S8C%5JY4,r&#W7U32XC2*$^5E",
    "AREA_ID": "Y@*9Ja.K17R%6^6R",
    "SUB_AREA_ID": "%.V@.O4H16rZ46A)",
    "PRINTER_HIST_SEQ": 2,
    "SPI_HIST_SEQ": 3,
    "REFLOW_HIST_SEQ": 7
  }
}
	`

func main() {
	opts, err := config.ParseClientFlags()
	if err != nil {
		log.Fatalln(err)
	}

	var conn *grpc.ClientConn

	if opts.Cert != "" {
		tlsCredentials, err := loadTLSCredentials(opts.Cert)
		if err != nil {
			log.Fatal("cannot load TLS credentials: ", err)
		}

		conn, err = grpc.Dial(opts.Target, grpc.WithTransportCredentials(tlsCredentials), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
	} else {
		conn, err = grpc.Dial(opts.Target, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
	}
	defer conn.Close()
	c := pb.NewOrderServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var pbd structpb.Struct
	if err := protojson.Unmarshal([]byte(jsonData), &pbd); err != nil {
		log.Panicln(err)
	}

	req := &pb.OrderRequest{
		Name:    opts.Name,
		Context: &pbd,
		Age:     123,
	}
	r, err := c.SayOrder(ctx, req)
	if err != nil {
		log.Panicf("could not greet: %v", err)
	}
	log.Printf("Order: %s, %+v, %d", r.GetName(), r.GetContext(), r.GetAge())
}

func loadTLSCredentials(cert string) (credentials.TransportCredentials, error) {
	pemServerCA, err := ioutil.ReadFile(cert)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	cfg := &tls.Config{
		RootCAs:    certPool,
		MinVersion: tls.VersionTLS12,
	}

	return credentials.NewTLS(cfg), nil
}
