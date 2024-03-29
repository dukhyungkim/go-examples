package config

import "github.com/jessevdk/go-flags"

type Options struct {
	ConfigPath string `long:"config" default:"config.yml" description:"path to config file"`
}

func ParseFlags() (*Options, error) {
	var options Options
	parser := flags.NewParser(&options, flags.Default)
	if _, err := parser.Parse(); err != nil {
		return nil, err
	}
	return &options, nil
}
