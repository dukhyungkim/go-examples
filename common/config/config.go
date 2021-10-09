package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Nats    *Nats    `yaml:"nats"`
	Redis   *Redis   `yaml:"redis"`
	MongoDB *MongoDB `yaml:"mongo_db"`
	Etcd    *Etcd    `yaml:"etcd"`
}

type Nats struct {
	Servers  []string `yaml:"servers"`
	Group    string   `yaml:"group"`
	Subject  string   `yaml:"subject"`
	Username string   `yaml:"username"`
	Password string   `yaml:"password"`
}

type Redis struct {
	Server   string `yaml:"server"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type MongoDB struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Etcd struct {
	Endpoints []string `yaml:"endpoints"`
}

func NewConfig(configPath string) (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
