package config

import "github.com/jessevdk/go-flags"

type ServerOptions struct {
	GRPCPort   int    `long:"grpc" default:"8000" description:"grpc server pot"`
	HTTPPort   int    `long:"http" default:"9000" description:"grpc-gateway server pot"`
	Cert       string `long:"cert" default:"" description:"certification path"`
	PrivateKey string `long:"key" default:"" description:"private key path"`
}

type ClientOptions struct {
	Target string `long:"target" short:"t" default:"localhost:8000" description:"target rpc server address"`
	Name   string `long:"name" short:"n" default:"world" description:"sending name"`
	Cert   string `long:"cert" default:"" description:"certification path"`
}

func ParseServerFlags() (*ServerOptions, error) {
	var options ServerOptions
	parser := flags.NewParser(&options, flags.Default)
	if _, err := parser.Parse(); err != nil {
		return nil, err
	}
	return &options, nil
}

func ParseClientFlags() (*ClientOptions, error) {
	var options ClientOptions
	parser := flags.NewParser(&options, flags.Default)
	if _, err := parser.Parse(); err != nil {
		return nil, err
	}
	return &options, nil
}
