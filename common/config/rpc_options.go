package config

import "github.com/jessevdk/go-flags"

type ServerOptions struct {
	Port int `long:"port" short:"p" default:"50051" description:"rpc server pot"`
}

type ClientOptions struct {
	Target string `long:"target" short:"t" default:"localhost:50051" description:"target rpc server address"`
	Name   string `long:"name" short:"n" default:"world" description:"sending name"`
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
