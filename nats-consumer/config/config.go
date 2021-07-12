package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Consumer Nats `yaml:"consumer"`
}

type Nats struct {
	Servers  []string `yaml:"servers"`
	Group    string   `yaml:"group"`
	Subject  string   `yaml:"subject"`
	Username string   `yaml:"username"`
	Password string   `yaml:"password"`
}

func NewConfig(configPath string) (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
