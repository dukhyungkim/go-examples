package config

import "github.com/jessevdk/go-flags"

type RPCOptions struct {
	Port int `long:"port" short:"p" default:"50051" description:"rpc server pot"`
}

func ParseRPCFlags() (*RPCOptions, error) {
	var options RPCOptions
	parser := flags.NewParser(&options, flags.Default)
	if _, err := parser.Parse(); err != nil {
		return nil, err
	}
	return &options, nil
}
