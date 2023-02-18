package main

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jessevdk/go-flags"
)

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

type Config struct {
	Redis *Redis `yaml:"redis"`
}

type Redis struct {
	Server   string `yaml:"server"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

func NewConfig(configPath string) (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
