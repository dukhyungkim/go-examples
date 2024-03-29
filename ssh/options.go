package main

import "github.com/jessevdk/go-flags"

type Options struct {
	Server string  `long:"server" description:"ssh server"`
	User   string  `long:"user" description:"ssh user"`
	Key    *string `long:"key" description:"ssh key path"`
}

func ParseFlags() (*Options, error) {
	var options Options
	parser := flags.NewParser(&options, flags.Default)
	if _, err := parser.Parse(); err != nil {
		return nil, err
	}
	return &options, nil
}
