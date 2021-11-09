package main

import (
	"fmt"
	"github.com/spf13/viper"
	"go-examples/common/config"
	"log"
)

func main() {
	var cfg config.Config
	readConfigFromFile(&cfg)

	log.Printf("%+v", cfg.Nats)
	log.Printf("%+v", cfg.Redis)
	log.Printf("%+v", cfg.MongoDB)
	log.Printf("%+v", cfg.Etcd)

	for _, key := range viper.AllKeys() {
		fmt.Printf("key: %s, val: %v\n", key, viper.Get(key))
	}
}

func readConfigFromFile(cfg *config.Config) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.SetEnvPrefix("my")
	viper.AutomaticEnv()

	log.Println("mongo-port", viper.GetInt("mongo.port"))

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	if err := viper.Unmarshal(cfg); err != nil {
		panic(fmt.Errorf("fatal error unmarshal config: %w", err))
	}
}
