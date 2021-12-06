package main

import "github.com/jessevdk/go-flags"

type Options struct {
	ImageName string `long:"image" description:"image name to management"`
	MaxCount  int    `long:"max" description:"max number of keeping"`
}

func ParseFlags() (*Options, error) {
	var options Options
	parser := flags.NewParser(&options, flags.Default)
	if _, err := parser.Parse(); err != nil {
		return nil, err
	}
	return &options, nil
}
