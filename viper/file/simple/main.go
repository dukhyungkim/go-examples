package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	readConfigFromFile()
	for _, key := range viper.AllKeys() {
		fmt.Printf("key: %s, val: %v\n", key, viper.Get(key))
	}
}

func readConfigFromFile() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}
