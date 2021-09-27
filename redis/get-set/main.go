package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go-examples/common/config"
	"log"
)

const (
	key           = "my_test_key"
	value         = "my_test_value"
	nonExistedKey = "non_existed_key"
)

var ctx = context.Background()

func main() {
	opts, err := config.ParseFlags()
	if err != nil {
		log.Fatalln(err)
	}

	cfg, err := config.NewConfig(opts.ConfigPath)
	if err != nil {
		log.Fatalf("Cannot access config: %v\n", err)
	}

	cache := NewCache(cfg.Redis)
	if err := cache.SetValue(key, value); err != nil {
		log.Fatalln(err)
	}

	val, err := cache.GetValue(key)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(key, val)

	val2, err := cache.GetValue(nonExistedKey)
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println(nonExistedKey, val2)
	}
}
