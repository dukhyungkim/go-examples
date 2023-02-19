package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jessevdk/go-flags"
)

type DBOptions struct {
	ConfigPath string `long:"config" default:"config.yml" description:"path to config file"`

	MariaDB   bool `long:"mariadb" description:"use MariaDB"`
	Postgres  bool `long:"postgres" description:"use PostgreSQL"`
	Migration bool `long:"migration" description:"auto migration"`
}

func ParseDBFlags() (*DBOptions, error) {
	var options DBOptions
	parser := flags.NewParser(&options, flags.Default)
	if _, err := parser.Parse(); err != nil {
		return nil, err
	}
	return &options, nil
}

type Config struct {
	Redis         *Redis         `yaml:"redis"`
	RDB           *RDB           `yaml:"rdb"`
	MongoDB       *MongoDB       `yaml:"mongo_db"`
	Etcd          *Etcd          `yaml:"etcd"`
	Harbor        *Harbor        `yaml:"harbor"`
	Elasticsearch *Elasticsearch `yaml:"elasticsearch"`
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

type Elasticsearch struct {
	Addresses []string `yaml:"addresses"`
	Username  string   `yaml:"username"`
	Password  string   `yaml:"password"`
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
