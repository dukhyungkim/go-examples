package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Nats    *Nats    `yaml:"nats"`
	Redis   *Redis   `yaml:"redis"`
	RDB     *RDB     `yaml:"rdb"`
	MongoDB *MongoDB `yaml:"mongo_db"`
	Etcd    *Etcd    `yaml:"etcd"`
	Harbor  *Harbor  `yaml:"harbor"`
}

type Nats struct {
	Servers  []string `yaml:"servers"`
	Group    string   `yaml:"group"`
	Subject  string   `yaml:"subject"`
	Username string   `yaml:"username"`
	Password string   `yaml:"password"`
}

type RDB struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	TimeZone string `yaml:"time_zone"`
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
	Username  string   `yaml:"username"`
	Password  string   `yaml:"password"`
}

type Harbor struct {
	APIHost  string `yaml:"api_host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func NewConfig(configPath string) (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
