package main

import (
	"fmt"
	"github.com/spf13/viper"
	"go-examples/common/config"
	"log"
)

func main() {
	var cfg config.Config
	readConfig(&cfg)

	log.Printf("%+v", cfg.Nats)
	log.Printf("%+v", cfg.Redis)
	log.Printf("%+v", cfg.MongoDB)
	log.Printf("%+v", cfg.Etcd)
}

func readConfig(cfg *config.Config) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	if err := viper.Unmarshal(cfg); err != nil {
		panic(fmt.Errorf("fatal error unmarshal config: %w", err))
	}
}
